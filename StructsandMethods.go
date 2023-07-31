package main

import (
	"fmt"
)

type ProfilePicture struct {
	ImageName string
	ImagePath string
}

// Nested structs
type SlackProfile struct {
	Name         string
	Username     string
	Designation  string
	ContactNumber string
	ProfilePicture
}

func (profile *SlackProfile) UpdateProfile(name, username, designation, contactNumber, imageName, imagePath string) {
	profile.Name = name
	profile.Username = username
	profile.Designation = designation
	profile.ContactNumber = contactNumber
	profile.ProfilePicture.ImageName = imageName
	profile.ProfilePicture.ImagePath = imagePath
}

func createSlackProfile() SlackProfile {
	var name, username, designation, contactNumber, imageName, imagePath string

	fmt.Println("Enter your Slack profile details:")
	fmt.Print("Name: ")
	fmt.Scanln(&name)

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Designation: ")
	fmt.Scanln(&designation)

	fmt.Print("Contact Number: ")
	fmt.Scanln(&contactNumber)

	fmt.Println("Enter your profile picture details:")
	fmt.Print("Image Name: ")
	fmt.Scanln(&imageName)

	fmt.Print("Image Path: ")
	fmt.Scanln(&imagePath)

	profile := SlackProfile{
		Name:         name,
		Username:     username,
		Designation:  designation,
		ContactNumber: contactNumber,
		ProfilePicture: ProfilePicture{
			ImageName: imageName,
			ImagePath: imagePath,
		},
	}
	return profile
}

func main() {
	myProfile := createSlackProfile()

	fmt.Println("\nSlack Profile")
	fmt.Println("Name:", myProfile.Name)
	fmt.Println("Username:", myProfile.Username)
	fmt.Println("Designation:", myProfile.Designation)
	fmt.Println("Contact Number:", myProfile.ContactNumber)
	fmt.Println("\nProfile Picture")
    // Promoted fields
	fmt.Println("Image Name:", myProfile.ImageName)
	fmt.Println("Image Path:", myProfile.ImagePath)

	// Update the profile
	myProfile.UpdateProfile("New Name", "new_username", "New Designation", "+9876543210", "new_profile_pic.png", "/new/path/to/profile/picture")

	fmt.Println("\nAfter Updating Slack Profile")
	fmt.Println("Name:", myProfile.Name)
	fmt.Println("Username:", myProfile.Username)
	fmt.Println("Designation:", myProfile.Designation)
	fmt.Println("Contact Number:", myProfile.ContactNumber)
	fmt.Println("\nProfile Picture")
	fmt.Println("Image Name:", myProfile.ProfilePicture.ImageName)
	fmt.Println("Image Path:", myProfile.ProfilePicture.ImagePath)
}
