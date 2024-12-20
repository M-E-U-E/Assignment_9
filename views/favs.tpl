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
    <script src="/static/js/favs.js?v={{ .UnixTimestamp }}"></script>
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
                  <span>Voting</span>
                </v-btn>
                <v-btn @click="navigateTo('/breeds')" class="nav-btn">
                  <v-icon left>mdi-magnify</v-icon> Breeds
                </v-btn>
                <v-btn @click="navigateTo('/favs')" class="nav-btn active">
                  <v-icon left>mdi-heart</v-icon> Favs
                </v-btn>
              </v-col>
            </v-row>

            <!-- Toggle View -->
            <v-row justify="center" class="mt-2">
              <v-btn @click="toggleView('grid')" class="nav-btn">
                <v-icon>mdi-view-grid</v-icon>
              </v-btn>
              <v-btn @click="toggleView('list')" class="nav-btn">
                <v-icon>mdi-view-list</v-icon>
              </v-btn>
            </v-row>

            <!-- Favorites Display -->
            <v-row justify="center" class="mt-2" :class="{ 'grid-view': view === 'grid', 'list-view': view === 'list' }">
              <v-col cols="12" md="4" lg="3" v-for="(fav, index) in favorites" :key="index">
                <v-card>
                  <v-img :src="fav.url" height="200px"></v-img>
                  <v-card-text>
                    <p>Favorite ID: {{ fav.id }}</p>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-content>
      </v-app>
    </div>
  </body>
</html>
