package models

type Users struct {
	User []Users
}

type User struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"creation_date"`
}

func AddUser(user User) error {
	tx := Db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", &user)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func AddUserNew(user User) error {
	_, err := Db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

func GetUserByName(name string) (users []User, err error) {
	err = Db.Select(&users, "SELECT * FROM users WHERE name=?", name)
	if len(users) != 0 {
		return users, err
	} else {
		return nil, err
	}
}

func GetUserByEmail(email string) (users []User, err error) {
	err = Db.Select(&users, "SELECT * FROM users WHERE email=?", email)
	return users, err
}
