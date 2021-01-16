package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hashed), nil
}

func CompareHashAndPassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func validateToken(tokenString string) (bool, error) {
	var secret = []byte("18O6PX88jh7TVGSVeoZqCS37Uh26lbF_OKsuIMq2zlBhEya8q66HHdQtjZAZE0oPDKPJHOqF79PkiNSgtdY35JdRY9SWeAYteYIGHc1iCqR-8tX5_BcKbA4VpACbvmQ8oXT0sLl1hqPgq-gBF-y2tCc_UzTmRRdGRIubZMGAj_qh93kGrEih7Hr2hrTmNNAXa_UIuIpxG4gm_6Dlq8WdvcKc3TSI7pRwM4XSa5QlI4gt29KzfQe6u-BA4_-VodPKUqkHP7Ya5S967615bc73-EdF-uwdsDTIVbLIxzVzLfED0tkOawQxA3AjNR-Yr7R2EWEsrpic_bq3Uvfb9ImilQ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return false, err
	}
	if token != nil {
		if token.Valid {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}
