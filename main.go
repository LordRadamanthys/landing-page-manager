package main

import (
	// "github.com/LordRadamanthys/landing-page-manager/app"
	// "github.com/LordRadamanthys/landing-page-manager/domain/brokers"

	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"github.com/LordRadamanthys/landing-page-manager/services"
	// "github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	// "github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
)

func main() {
	// app.StartApplication()
	broker := brokers.Brokers{

		Id_broker:                2,
		Name:                     "Tatine de Lima",
		Field_broker_title:       "Title template 2",
		Field_broker_description: "desc 2",
		Field_broker_imageCenter: "https://static-cdn.jtvnw.net/jtv_user_pictures/c5327992-ba31-44d6-9cfc-557aedb23237-profile_image-70x70.png",
		Field_broker_imageLeft:   "https://blog.logrocket.com/wp-content/uploads/2021/10/How-to-use-MongoDB-Go.jpg",
	}

	// lp := landing_page.LandingPage{
	// 	Id:                "620b284cfbc8aa595f6e66b9",
	// 	Title:             "Natal",
	// 	Template:          "1",
	// 	Id_mkt:            1,
	// 	Id_categoria:      "1",
	// 	Active:            true,
	// 	Date_Created:      time.Now(),
	// 	Field_title:       "Feliz natal",
	// 	Field_description: "Feliz natal",
	// 	Field_image:       "image",
	// 	// BrokersList:       nil,
	// }

	// services.LandingPageService.Insert(lp)
	// services.BrokerService.GetTemplates("1")
	services.BrokerService.Insert("620fd3ab2f4a16045649e324", broker)
}
