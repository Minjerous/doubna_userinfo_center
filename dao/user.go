package dao

import "doubna_userinfo/model"

func QueryByEmail(email string) (model.User, error) {
	user := model.User{}

	sqlStr := "SELECT id,username,password,email,phone,salt,avatar,domain_name,habitat,hometown,birthday,statement,followers,followings FROM users WHERE emial= ? "
	Stmt, err := DB.Prepare(sqlStr)
	defer Stmt.Close()

	if err != nil {
		return user, err
	}

	row := Stmt.QueryRow(email)
	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Salt, &user.Avatar, &user.DomainName, &user.Habitat, &user.Username, &user.Birthday, &user.Statement, &user.Followers, &user.Followings)
	if err != nil {
		return user, err
	}

	return user, nil
}


func QueryByPhone(phone string) (model.User, error) {
	user := model.User{}
	rs := Db.Where(&model.User{Phone: phone}).Find(&user)
	return user, rs.Error
}