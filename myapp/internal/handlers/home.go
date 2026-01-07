package handlers

import (
  "net/http"
  "math/rand"
  "time"
)

func Home(w http.ResponseWriter, r *http.Request) {
    names := []string{"Joe", "Bob", "Jessica", "Nathan", "Becca"}
    data := map[string]string{"Name": randomName(names)}
    renderTemplate(w, "home.html", data)
}

func About(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.html", nil)
}

func randomName(names []string) string{
  if len(names) == 0 {
    return ""
  }
  rand.Seed(time.Now().UnixNano())
  index := rand.Intn(len(names))
  return names[index]
}
