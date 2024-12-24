<!-- Template (index.tpl) -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Favorites</title>
    <link href="https://fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900|Material+Icons" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/vuetify@2.x/dist/vuetify.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/favs.css?v={{ .UnixTimestamp }}">
    <script src="https://cdn.jsdelivr.net/npm/vue@2.x/dist/vue.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/vuetify@2.x/dist/vuetify.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
  </head>
  <body>
    <div id="app">
      <v-app>
        <v-content>
          <v-container>
            <!-- Navigation -->
            <v-row justify="center" class="mt-4">
              <v-col cols="12" md="6" class="nav-container">
                <v-btn @click="navigateTo('/voting')" class="nav-btn">
                  <v-icon class="icon-style">mdi-arrow-up</v-icon>
                  <v-icon class="icon-style">mdi-arrow-down</v-icon>
                  Voting
                </v-btn>
                <v-btn @click="navigateTo('/breeds')" class="nav-btn">
                  <v-icon>mdi-magnify</v-icon> Breeds
                </v-btn>
                <v-btn @click="navigateTo('/favs')" class="nav-btn active">
                  <v-icon>mdi-heart</v-icon> Favs
                </v-btn>
              </v-col>
            </v-row>

            <!-- View Toggle Buttons -->
            <v-row justify="center" class="mt-4">
              <v-col cols="12" md="6" class="view-toggle-container">
                <v-btn-toggle v-model="viewMode" mandatory>
                  <v-btn value="grid">
                    <v-icon>mdi-grid</v-icon>
                  </v-btn>
                  <v-btn value="sequence">
                    <v-icon>mdi-format-list-bulleted</v-icon>
                  </v-btn>
                </v-btn-toggle>
              </v-col>
            </v-row>

            <!-- Dynamic Image Display -->
            <div :class="['image-display', viewMode]">
              <div v-for="(image, index) in favourites" 
                   :key="index" 
                   class="image-container">
                <img :src="image.image.url" 
                     alt="Favorite Image" 
                     class="fav-img" />
              </div>
            </div>
          </v-container>
        </v-content>
      </v-app>
    </div>
    <script src="/static/js/favs.js?v={{ .UnixTimestamp }}"></script>
  </body>
</html>