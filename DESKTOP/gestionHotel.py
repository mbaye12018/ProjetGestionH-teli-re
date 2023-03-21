# -*- coding: utf-8 -*-
import gi
gi.require_version('Gtk', '3.0')
from gi.repository import Gtk

class MainWindow(Gtk.Window):
    

    def __init__(self):
        Gtk.Window.__init__(self, title="Hotel GLSIB")

        # Création de la barre d'en-tête
        headerbar = Gtk.HeaderBar()
        headerbar.set_show_close_button(True)
        headerbar.props.title = "Hotel GLSIB"
        self.set_titlebar(headerbar)
        # Définition de l'image d'arrière-plan
        image = Gtk.Image()
        image.set_from_file("C:\ht/cordoba.JPG")
        self.add(image)

        # Création de la boîte principale
        self.box = Gtk.Box(spacing=6)
        self.add(self.box)

        # Création du bouton "Client"
        self.button_client = Gtk.Button(label="Client")
        self.button_client.connect("clicked", self.on_client_clicked)
        headerbar.pack_start(self.button_client)

        # Création du bouton "Chambre"
        self.button_chambre = Gtk.Button(label="Chambre")
        self.button_chambre.connect("clicked", self.on_chambre_clicked)
        headerbar.pack_start(self.button_chambre)

        # Création du bouton "Réservation"
        self.button_Reservation = Gtk.Button(label="Reservation")
        self.button_Reservation.connect("clicked", self.on_reservation_clicked)
        headerbar.pack_start(self.button_Reservation)

        # Création du bouton "Facture"
        self.button_facture = Gtk.Button(label="Facture")
        headerbar.pack_end(self.button_facture)

        # Création du bouton "Statistiques"
        self.button_statistiques = Gtk.Button(label="Statistiques")
        headerbar.pack_end(self.button_statistiques)
        

    def on_client_clicked(self, widget):
        # Ouvrir une nouvelle fenêtre pour le formulaire du client
        win = Gtk.Window(title="Formulaire client")
        win.connect("destroy", Gtk.main_quit)
        win.set_default_size(200, 100)

        # Création de la boîte pour le formulaire
        box = Gtk.Box(spacing=6)
        win.add(box)

        # Ajout des champs pour le formulaire
        prenom_label = Gtk.Label("Prénom")
        box.pack_start(prenom_label, True, True, 0)
        prenom_entry = Gtk.Entry()
        box.pack_start(prenom_entry, True, True, 0)

        nom_label = Gtk.Label("Nom")
        box.pack_start(nom_label, True, True, 0)
        nom_entry = Gtk.Entry()
        box.pack_start(nom_entry, True, True, 0)

        adresse_label = Gtk.Label("Adresse")
        box.pack_start(adresse_label, True, True, 0)
        adresse_entry = Gtk.Entry()
        box.pack_start(adresse_entry, True, True, 0)

        telephone_label = Gtk.Label("Téléphone")
        box.pack_start(telephone_label, True, True, 0)
        telephone_entry = Gtk.Entry()
        box.pack_start(telephone_entry, True, True, 0)

        # Ajout du bouton "Enregistrer"
        button_enregistrer = Gtk.Button(label="Enregistrer")
        box.pack_start(button_enregistrer, True, True, 0)
        win.show_all()

    def on_chambre_clicked(self, widget):
        # Ouvrir une nouvelle fenêtre pour les options de chambre
        win = Gtk.Window(title="Options de chambre")
        win.connect("destroy", Gtk.main_quit)
        win.set_default_size(200, 100)

        # Création de la boîte pour les options
        box = Gtk.Box(spacing=6)
        win.add(box)

        # Ajout des boutons pour les différentes options
        button_liste_chambres = Gtk.Button(label="Liste des chambres")
        box.pack_start(button_liste_chambres, True, True, 0)

        button_chambres_occupees = Gtk.Button(label="Liste des chambres occupées")
        box.pack_start(button_chambres_occupees, True, True, 0)

        button_chambres_reservees = Gtk.Button(label="Liste des chambres réservées")
        box.pack_start(button_chambres_reservees, True, True, 0) 

        win.show_all()

    def on_reservation_clicked(self, widget):
        # Ouvrir une nouvelle fenêtre pour le formulaire du client
        win = Gtk.Window(title="Formulaire reservation")
        win.connect("destroy", Gtk.main_quit)
        win.set_default_size(200, 100)

        # Création de la boîte pour le formulaire
        box = Gtk.Box(spacing=6)
        win.add(box)

        # Ajout des champs pour le formulaire
        Date_arrivee_label = Gtk.Label("Date_arrivee")
        box.pack_start(Date_arrivee_label, True, True, 0)
        Date_arrivee_entry = Gtk.Entry()
        box.pack_start(Date_arrivee_entry, True, True, 0)

        Date_depart_label = Gtk.Label("Date_depart")
        box.pack_start(Date_depart_label, True, True, 0)
        Date_depart_entry = Gtk.Entry()
        box.pack_start(Date_depart_entry, True, True, 0)

        Type_tarif_label = Gtk.Label("Type_tarif")
        box.pack_start(Type_tarif_label, True, True, 0)
        Type_tarif_entry = Gtk.Entry()
        box.pack_start(Type_tarif_entry, True, True, 0)
        # Ajout du bouton "Enregistrer"
        button_enregistrer = Gtk.Button(label="Enregistrer")
        box.pack_start(button_enregistrer, True, True, 0)
        win.show_all()

      
       

win = MainWindow()
win.connect("destroy", Gtk.main_quit)
win.show_all()
Gtk.main()
