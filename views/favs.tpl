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
    <script src="/static/js/favs.js?v=1.0"></script>
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

            <!-- Favorites Gallery -->
            <v-row justify="center" class="mt-4">
              <v-col cols="12" md="6">
                <div v-if="favourites.length > 0">
                  <v-row>
                    <v-col cols="12" md="4" v-for="(favorite, index) in favourites" :key="favorite.id">
                      <v-img :src="favorite.image.url" class="fav-img"></v-img>
                    </v-col>
                  </v-row>
                </div>
                <p v-else><h1>Loading...</h1></p>
              </v-col>
            </v-row>

            <!-- Pagination Controls -->
            <v-row justify="center" class="mt-4">
              <v-pagination
                v-if="pagination_count > limit"
                v-model="page"
                :length="Math.ceil(pagination_count / limit)"
                @input="changePage"
              ></v-pagination>
            </v-row>
          </v-container>
        </v-content>
      </v-app>
    </div>
    <script src="/static/js/favs.js?v={{ .UnixTimestamp }}"></script>
  </body>
</html>
