package main

import "fmt"

// Purpose:
//		- Encapsulates an object's construction process along with
//		  specifying the various parts of a complex API
//		- Enables flexible creation of an object that can have
//		  different representations
//		- Increases code readability for complex types
//
// Scenarios:
//		- Objects that have complex APIs, multiple constructor options,
//		  and several possible representations

type NotificationBuilder struct {
	Title            string
	SubTitle         string
	Message          string
	Image            string
	Icon             string
	Priority         int
	NotificationType string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notificationType string) {
	nb.NotificationType = notificationType
}

// The Build method returns a fully finished Notification object
func (nb *NotificationBuilder) Build() (*Notification, error) {
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("you need to specify a subtitle when using an icon")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority must be 0 to 5")
	}

	return &Notification{
		title:            nb.Title,
		subtitle:         nb.SubTitle,
		message:          nb.Message,
		image:            nb.Image,
		icon:             nb.Icon,
		priority:         nb.Priority,
		notificationType: nb.NotificationType,
	}, nil
}
