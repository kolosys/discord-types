package v10

import (
	"encoding/json"
	"fmt"
	"log"
)

// Example demonstrates how to handle polymorphic message components
func ExampleHandleMessageComponents() {
	// Example JSON response from Discord API containing mixed component types
	jsonResponse := `{
		"id": "123456789012345678",
		"channel_id": "987654321098765432",
		"content": "Choose an option:",
		"components": [
			{
				"type": 1,
				"components": [
					{
						"type": 2,
						"style": 1,
						"label": "Primary Button",
						"custom_id": "primary_btn"
					},
					{
						"type": 2,
						"style": 3,
						"label": "Success Button",
						"custom_id": "success_btn"
					}
				]
			},
			{
				"type": 1,
				"components": [
					{
						"type": 3,
						"custom_id": "select_menu",
						"placeholder": "Choose an option",
						"options": [
							{
								"label": "Option 1",
								"value": "opt1"
							},
							{
								"label": "Option 2", 
								"value": "opt2"
							}
						]
					}
				]
			}
		]
	}`

	// Unmarshal into Message struct
	var message Message
	if err := json.Unmarshal([]byte(jsonResponse), &message); err != nil {
		log.Fatal("Failed to unmarshal message:", err)
	}

	// Process each action row
	for i, actionRow := range message.Components {
		fmt.Printf("Action Row %d:\n", i+1)

		// Process each component in the action row
		for j, component := range actionRow.Components {
			fmt.Printf("  Component %d (Type: %d): ", j+1, component.GetType())

			// Handle different component types
			switch {
			case ComponentChecker.IsButton(component):
				if button, ok := ComponentChecker.AsButton(component); ok {
					fmt.Printf("Button - Label: %s, Style: %d, CustomID: %s\n",
						getStringValue(button.Label),
						button.Style,
						getStringValue(button.CustomID))
				}

			case ComponentChecker.IsSelectMenu(component):
				if selectMenu, ok := ComponentChecker.AsSelectMenu(component); ok {
					fmt.Printf("Select Menu - CustomID: %s, Placeholder: %s, Options: %d\n",
						selectMenu.CustomID,
						getStringValue(selectMenu.Placeholder),
						len(selectMenu.Options))
				}

			case ComponentChecker.IsTextInput(component):
				if textInput, ok := ComponentChecker.AsTextInput(component); ok {
					fmt.Printf("Text Input - Label: %s, CustomID: %s, Style: %d\n",
						textInput.Label,
						textInput.CustomID,
						textInput.Style)
				}

			default:
				fmt.Printf("Unknown component type\n")
			}
		}
		fmt.Println()
	}
}

// Helper function to safely get string values from pointers
func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// ExampleCreateComponents demonstrates how to create components programmatically
func ExampleCreateComponents() {
	// Create a button component
	button := ButtonComponent{
		Type:     ComponentTypeButton,
		Style:    ButtonStylePrimary,
		Label:    stringPtr("Click Me!"),
		CustomID: stringPtr("my_button"),
	}

	// Create a select menu component
	selectMenu := SelectMenuComponent{
		Type:     ComponentTypeStringSelect,
		CustomID: "my_select",
		Options: []SelectOption{
			{
				Label: "Option 1",
				Value: "opt1",
			},
			{
				Label: "Option 2",
				Value: "opt2",
			},
		},
		Placeholder: stringPtr("Choose an option"),
	}

	// Create an action row containing both components
	actionRow := ActionRowComponent{
		Type: ComponentTypeActionRow,
		Components: []MessageComponent{
			button,
			selectMenu,
		},
	}

	// Marshal to JSON to see the result
	jsonData, err := json.MarshalIndent(actionRow, "", "  ")
	if err != nil {
		log.Fatal("Failed to marshal components:", err)
	}

	fmt.Println("Created components JSON:")
	fmt.Println(string(jsonData))
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

// ExampleInteractionHandling shows how to handle interaction responses with components
func ExampleInteractionHandling() {
	// Example interaction response with components
	response := InteractionResponseCallbackData{
		Content: stringPtr("Please choose an option:"),
		Components: []ActionRowComponent{
			{
				Type: ComponentTypeActionRow,
				Components: []MessageComponent{
					ButtonComponent{
						Type:     ComponentTypeButton,
						Style:    ButtonStylePrimary,
						Label:    stringPtr("Yes"),
						CustomID: stringPtr("confirm_yes"),
					},
					ButtonComponent{
						Type:     ComponentTypeButton,
						Style:    ButtonStyleSecondary,
						Label:    stringPtr("No"),
						CustomID: stringPtr("confirm_no"),
					},
				},
			},
		},
	}

	// Marshal the response
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal("Failed to marshal interaction response:", err)
	}

	fmt.Println("Interaction response with components:")
	fmt.Println(string(jsonData))
}
