package permissions

import (
	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Permissions struct {
	WriteMember      bool
	DeleteMember     bool
	WriteEnterprise  bool
	DeleteEnterprise bool
	WriteParent      bool
	DeleteParent     bool
	WritePayment     bool
	DeletePayment    bool
}

func GetAdmin(user string) (bool, error) {
	row := database.DB.QueryRow("SELECT Admin FROM UserTable WHERE Username = ?", user)
	var admin bool
	err := row.Scan(&admin)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return false, er.ScanError
	}
	return admin, nil
}

func GetPermissions(user string) (Permissions, error) {
	row := database.DB.QueryRow("SELECT WriteMember, DeleteMember, WriteEnterprise, DeleteEnterprise, WriteParent, DeleteParent, WritePayment, DeletePayment FROM UserTable WHERE Username = ?", user)
	p := Permissions{}

	err := row.Scan(&p.WriteMember, &p.DeleteMember, &p.WriteEnterprise, &p.DeleteEnterprise, &p.WriteParent, &p.DeleteParent, &p.WritePayment, &p.DeletePayment)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return Permissions{}, er.ScanError
	}

	return p, nil

}

func VerifyAdmin(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["admin"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyWriteMember(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["writeMember"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyDeleteMember(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["deleteMember"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyWriteEnterprise(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyDeleteEnterprise(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyWriteParent(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["writeParent"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyDeleteParent(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["deleteParent"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyWritePayment(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["writePayment"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyDeletePayment(c *fiber.Ctx) error {
	permission := c.Locals("claims").(jwt.MapClaims)["deletePayment"]
	if permission == true {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}
