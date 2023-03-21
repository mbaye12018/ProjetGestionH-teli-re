package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

type Client struct {
	id_client        int
	Nom_client       string
	Prenom_client    string
	Adresse_client   string
	Telephone_client int
}
type Reservation struct {
	ID_reservation int
	Date_arrivee   time.Time
	Date_depart    time.Time
	Type_tarif     string // Ou un type enum, si votre base de données le permet
	ID_client      int
	Num_chambre    int
}

type Chambre struct {
	Num_chambre  int
	ID_etage     int
	ID_categorie int
	Statut       string // Ou un type enum, si votre base de données le permet
	Nom_hotel    string
}

type Hotel struct {
	Nom_hotel            string
	Nb_etages            int
	Nb_chambre_par_etage int
}

type Categorie struct {
	ID_categorie   int
	Nom_categorie  string // Ou un type enum, si votre base de données le permet
	Tarif_unitaire int
}

type Service struct {
	ID_service     int
	Nom_service    string // Ou un type enum, si votre base de données le permet
	ID_reservation int
	Nom_hotel      string
}

type Etage struct {
	ID_etage int
}

func main() {
	// les données de notre serveur proxy SQL
	server := "proxyserver"
	port := 6032
	user := "cluster_admin"
	password := "cluster_admin_password"
	database := "gestionhotel"

	// Créez une chaîne de connexion
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	// Ouvrez une connexion
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err.Error())
	}

	// Fermez la connexion lorsque vous avez terminé
	defer db.Close()

	// Gestion de l'API REST pour les clients
	http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		// Vérifiez la méthode HTTP
		if r.Method == "GET" {
			// Exécutez une requête SQL pour récupérer tous les clients
			rows, err := db.Query("SELECT * FROM clients")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			// Créez une slice pour stocker les clients
			clients := make([]Client, 0)

			// Parcourez les lignes de résultats et stockez les clients dans la slice
			for rows.Next() {
				client := Client{}
				err := rows.Scan(&client.id_client, &client.Nom_client, &client.Prenom_client, &client.Adresse_client, &client.Telephone_client)
				if err != nil {
					log.Fatal(err)
				}
				clients = append(clients, client)
			}

			// Encodez la slice de clients en JSON et envoyez la réponse
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(clients)
			if err != nil {
				log.Fatal(err)
			}
		} else if r.Method == "POST" {
			// Lecture des données envoyées par le client
			var client Client
			err := json.NewDecoder(r.Body).Decode(&client)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Exécutez une requête SQL pour insérer un nouveau client
			query := fmt.Sprintf("INSERT INTO clients(Nom_client, Prenom_client, Adresse_client, Telephone_client) VALUES ('%s', '%s', '%s', '%d')", client.Nom_client, client.Prenom_client, client.Adresse_client, client.Telephone_client)
			_, err = db.Exec(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Envoi de la réponse au client
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Le client a été ajouté avec succès.")
		} else if r.Method == "PUT" {
			// Lecture des données envoyées par le client
			var client Client
			err := json.NewDecoder(r.Body).Decode(&client)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Exécutez une requête SQL pour mettre à jour le client existant
			query := fmt.Sprintf("UPDATE clients SET Nom_client='%s', Prenom_client='%s', Adresse_client='%s', Telephone_client='%d' WHERE id_client=%d", client.Nom_client, client.Prenom_client, client.Adresse_client, client.Telephone_client, client.id_client)
			_, err = db.Exec(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Envoi de la réponse au client
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Le client a été mis à jour avec succès.")
		} else if r.Method == "DELETE" {
			// Lecture de l'ID du client à supprimer
			clientID := r.URL.Path[len("/clients/"):]

			// Exécutez une requête SQL pour supprimer le client correspondant à l'ID
			query := fmt.Sprintf("DELETE FROM clients WHERE id_client=%s", clientID)
			result, err := db.Exec(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Vérifiez si un client a été supprimé et envoyez la réponse
			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 0 {
				http.Error(w, "Le client n'existe pas.", http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Le client a été supprimé avec succès.")
		} else {
			// Méthode non autorisée
			http.Error(w, "Méthode non autorisée.", http.StatusMethodNotAllowed)
		}
	})

	// Démarrez le serveur
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Handler pour l'ajout d'une nouvelle réservation de client
	http.HandleFunc("/reservations", func(w http.ResponseWriter, r *http.Request) {
		// Vérifiez la méthode HTTP
		if r.Method == "POST" {
			// Lecture des données envoyées par le client
			var reservation Reservation
			err := json.NewDecoder(r.Body).Decode(&reservation)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Vérification de la disponibilité de la chambre
			rows, err := db.Query(fmt.Sprintf("SELECT * FROM chambres WHERE Num_chambre = %d AND Statut = 'libre'", reservation.Num_chambre))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			if !rows.Next() {
				http.Error(w, "La chambre n'est pas disponible", http.StatusBadRequest)
				return
			}

			// Exécutez une requête SQL pour insérer une nouvelle réservation
			query := fmt.Sprintf("INSERT INTO reservations(Date_arrivee, Date_depart, Type_tarif, ID_client, Num_chambre) VALUES ('%s', '%s', '%s', %d, %d)", reservation.Date_arrivee.Format("2006-01-02"), reservation.Date_depart.Format("2006-01-02"), reservation.Type_tarif, reservation.ID_client, reservation.Num_chambre)
			_, err = db.Exec(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Mettre à jour le statut de la chambre réservée
			query = fmt.Sprintf("UPDATE chambres SET Statut = 'occupé' WHERE Num_chambre = %d", reservation.Num_chambre)
			_, err = db.Exec(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Envoi de la réponse au client
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "La réservation a été ajoutée avec succès.")
		}
	})
	http.HandleFunc("/reservations", func(w http.ResponseWriter, r *http.Request) {
		// Vérifiez la méthode HTTP
		if r.Method == "GET" {
			// Exécutez une requête SQL pour récupérer toutes les réservations
			rows, err := db.Query("SELECT * FROM reservations")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			// Créez une slice pour stocker les réservations
			reservations := make([]Reservation, 0)

			// Parcourez les lignes de résultats et stockez les réservations dans la slice
			for rows.Next() {
				reservation := Reservation{}
				err := rows.Scan(&reservation.ID_reservation, &reservation.Date_arrivee, &reservation.Date_depart, &reservation.Type_tarif, &reservation.ID_client, &reservation.Num_chambre)
				if err != nil {
					log.Fatal(err)
				}
				reservations = append(reservations, reservation)
			}

			// Encodez la slice de réservations en JSON et envoyez la réponse
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(reservations)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// Si la méthode HTTP n'est pas GET, renvoyez une erreur "Méthode non autorisée"
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/reservations/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Vérifiez la méthode HTTP
		if r.Method == "PUT" {
			// Récupérez l'ID de la réservation à modifier à partir des paramètres de l'URL
			vars := mux.Vars(r)
			reservationID, err := strconv.Atoi(vars["id"])
			if err != nil {
				http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
				return
			}

			// Decodez la requête JSON dans une struct Reservation
			var reservation Reservation
			err = json.NewDecoder(r.Body).Decode(&reservation)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Exécutez une requête SQL pour mettre à jour la réservation avec l'ID spécifié
			_, err = db.Exec("UPDATE reservations SET Date_arrivee=?, Date_depart=?, Type_tarif=?, ID_client=?, Num_chambre=? WHERE ID_reservation=?", reservation.Date_arrivee, reservation.Date_depart, reservation.Type_tarif, reservation.ID_client, reservation.Num_chambre, reservationID)
			if err != nil {
				log.Fatal(err)
			}

			// Récupérez la réservation mise à jour de la base de données et encodez-la en JSON
			updatedReservation := Reservation{}
			err = db.QueryRow("SELECT * FROM reservations WHERE ID_reservation=?", reservationID).Scan(&updatedReservation.ID_reservation, &updatedReservation.Date_arrivee, &updatedReservation.Date_depart, &updatedReservation.Type_tarif, &updatedReservation.ID_client, &updatedReservation.Num_chambre)
			if err != nil {
				log.Fatal(err)
			}
			responseJSON, err := json.Marshal(updatedReservation)
			if err != nil {
				log.Fatal(err)
			}

			// Renvoyer la réservation mise à jour
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseJSON)
		} else {
			// Si la méthode HTTP n'est pas PUT, renvoyez une erreur "Méthode non autorisée"
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}
