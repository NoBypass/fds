package lifecycle

import "log"

func (b *Bot) Shutdown() {
	if *b.RemoveCommands {
		log.Println("Removing commands...")
		userID := b.Session.State.User.ID

		registeredCommands, err := b.Session.ApplicationCommands(userID, "")
		if err != nil {
			log.Fatalf("Could not fetch registered commands: %v", err)
		}

		for _, v := range registeredCommands {
			err := b.Session.ApplicationCommandDelete(userID, "", v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
			log.Printf("Deleted '%v' command", v.Name)
		}
	}

	log.Println("Gracefully shutting down.")
}
