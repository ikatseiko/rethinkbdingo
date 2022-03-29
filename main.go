package main

import "fmt"

func main() {
	ss := conect()
	if ss != nil {
		b := false

		//b = AddClient(ss, &Client{
		//	Name: "OI",
		//})

		clients := AllClients(ss)
		b = len(*clients) > 0
		if b {
			fmt.Println("clients:")
			for i, client := range *clients {
				fmt.Printf("  %d) %+v\n", i+1, client)
				if client.Name == "" {
					DelClient(ss, client.Id)
					fmt.Printf("  %d) DEL\n", i+1)

				}
			}
			fmt.Println("")
		}

		users := AllUsers(ss)
		b = len(*users) > 0
		if b {
			fmt.Println("users:")
			for i, user := range *users {
				fmt.Printf("  %d) %+v\n", i+1, user)
				if user.Status == UserNew {
					user.Status = UserActive
					UpdateUser(ss, &user)
					fmt.Printf("  %d) UPDATE\n", i+1)
					fmt.Printf("  %d) %+v\n", i+1, user)

				}
				// UserName
				if user.Login == "UserName" {
					DelUser(ss, user.Id)
					fmt.Printf("  %d) DEL\n", i+1)

				}
			}
			fmt.Println("")
		}

		//b = AddUser(ss, &User{
		//	Login:  "UserName",
		//	Status: UserNew,
		//})

		if b {
			fmt.Println("nice!")
		}
		if err := close(ss); err != nil {
			fmt.Println("ups")
		}
	}

}
