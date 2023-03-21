

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema gestionhotel
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema gestionhotel
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `gestionhotel` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `gestionhotel` ;

-- -----------------------------------------------------
-- Table `gestionhotel`.`categorie`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`categorie` (
  `Idcategorie` INT NOT NULL AUTO_INCREMENT,
  `Nom_categorie` ENUM('economique', 'standing', 'affaires') NOT NULL,
  `Tarif_unitaire` DECIMAL(10,0) NOT NULL,
  PRIMARY KEY (`Idcategorie`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `gestionhotel`.`etage`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`etage` (
  `Id_etage` INT NOT NULL,
  PRIMARY KEY (`Id_etage`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `gestionhotel`.`hotel`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`hotel` (
  `Nom_hotel` VARCHAR(45) NOT NULL,
  `NbEtages` INT NOT NULL,
  `NbChambresParEtage` INT NOT NULL,
  PRIMARY KEY (`Nom_hotel`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `gestionhotel`.`chambre`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`chambre` (
  `Num_chambre` INT NOT NULL AUTO_INCREMENT,
  `Id_etage` INT NOT NULL,
  `Idcategorie` INT NOT NULL,
  `Statut` ENUM('occupe', 'libre') NOT NULL,
  `Nom_hotel` VARCHAR(45) NULL,
  PRIMARY KEY (`Num_chambre`),
  INDEX `idcategorie_idx` (`Idcategorie` ASC) VISIBLE,
  INDEX `id_etage_idx` (`Id_etage` ASC, `Idcategorie` ASC) VISIBLE,
  INDEX `nom_hotel_idx` (`Nom_hotel` ASC) VISIBLE,
  CONSTRAINT `id_etage`
    FOREIGN KEY (`Id_etage`)
    REFERENCES `gestionhotel`.`etage` (`Id_etage`),
  CONSTRAINT `idcategorie`
    FOREIGN KEY (`Idcategorie`)
    REFERENCES `gestionhotel`.`categorie` (`Idcategorie`)
    ON DELETE SET NULL
    ON UPDATE SET NULL,
  CONSTRAINT `nom_hotel`
    FOREIGN KEY (`Nom_hotel`)
    REFERENCES `gestionhotel`.`hotel` (`Nom_hotel`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `gestionhotel`.`client`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`client` (
  `Id_client` INT NOT NULL AUTO_INCREMENT,
  `Nom_client` VARCHAR(45) NOT NULL,
  `Prenom_client` VARCHAR(45) NOT NULL,
  `Adresse_client` VARCHAR(45) NULL DEFAULT NULL,
  `Telephone_client` INT NULL DEFAULT NULL,
  PRIMARY KEY (`Id_client`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `gestionhotel`.`reservation`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`reservation` (
  `Id_reservation` INT NOT NULL AUTO_INCREMENT,
  `Date_arrivee` DATE NOT NULL,
  `Date_depart` DATE NOT NULL,
  `Type_tarif` ENUM('normal', 'groupe') NOT NULL,
  `Id_client` INT NULL DEFAULT NULL,
  `Num_chambre` INT NULL DEFAULT NULL,
  PRIMARY KEY (`Id_reservation`),
  INDEX `id_client_idx` (`Id_client` ASC) VISIBLE,
  INDEX `num_chambre_idx` (`Num_chambre` ASC) VISIBLE,
  CONSTRAINT `id_client`
    FOREIGN KEY (`Id_client`)
    REFERENCES `gestionhotel`.`client` (`Id_client`),
  CONSTRAINT `num_chambre`
    FOREIGN KEY (`Num_chambre`)
    REFERENCES `gestionhotel`.`chambre` (`Num_chambre`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `gestionhotel`.`service`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `gestionhotel`.`service` (
  `Id_service` INT NOT NULL AUTO_INCREMENT,
  `Nom_service` ENUM('ptit dejeuner', 'bar', 'telephone') NOT NULL,
  `Tarif_unitaire` DECIMAL(10,0) NOT NULL,
  `Id_reservation` INT NULL,
  `Nom_hotel` VARCHAR(45) NULL,
  PRIMARY KEY (`Id_service`),
  INDEX `id_reservation_idx` (`Id_reservation` ASC) VISIBLE,
  INDEX `nom_hotel_idx` (`Nom_hotel` ASC) VISIBLE,
  CONSTRAINT `id_reservation`
    FOREIGN KEY (`Id_reservation`)
    REFERENCES `gestionhotel`.`reservation` (`Id_reservation`)
    ON DELETE SET NULL
    ON UPDATE SET NULL,
  CONSTRAINT `nom_hotel`
    FOREIGN KEY (`Nom_hotel`)
    REFERENCES `gestionhotel`.`hotel` (`Nom_hotel`)
    ON DELETE SET NULL
    ON UPDATE SET NULL)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
