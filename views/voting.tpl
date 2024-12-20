<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Voting</title>
    <link href="https://fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900|Material+Icons" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/vuetify@2.x/dist/vuetify.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/voting.css?v={{ .UnixTimestamp }}">
    <script src="https://cdn.jsdelivr.net/npm/vue@2.x/dist/vue.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/vuetify@2.x/dist/vuetify.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <script src="/static/js/voting.js?v=1.0"></script>
  </head>
  <body>
    <div id="app">
      <v-app>
        <v-content>
          <v-container>
            <!-- Navigation -->
            <v-row justify="center" class="mt-4">
              <v-col cols="12" md="6" class="nav-container">
                <v-btn @click="navigateTo('/voting')" class="nav-btn" :class="{ active: activeTab === 'voting' }">
                  <v-icon class="icon-style">mdi-arrow-up</v-icon>
                  <v-icon class="icon-style">mdi-arrow-down</v-icon>
                  <span>Voting</span>
                </v-btn>
                <v-btn @click="navigateTo('/breeds')" class="nav-btn" :class="{ active: activeTab === 'breeds' }">
                  <v-icon left>mdi-magnify</v-icon> Breeds
                </v-btn>
                <v-btn @click="navigateTo('/favs')" class="nav-btn" :class="{ active: activeTab === 'favs' }">
                  <v-icon left>mdi-heart</v-icon> Favs
                </v-btn>
              </v-col>
            </v-row>

            <!-- Image Card -->
            <v-row justify="center" class="mt-2">
              <v-col cols="12" md="6">
                <v-card>
                  <!-- Cat Image -->
                  <v-img :src="catImage" height="400px"></v-img>

                  <!-- Action Buttons -->
                  <v-card-actions>
                    <v-btn icon @click="addToFavorites" class="fav-btn">
                        <v-icon color="red">mdi-heart</v-icon>
                    </v-btn>
                    <v-spacer></v-spacer>
                    <v-btn icon @click="fetchCatImage" class="upvote-btn">
                      <v-icon color="green">mdi-thumb-up</v-icon>
                    </v-btn>
                    <v-btn icon @click="fetchCatImage" class="downvote-btn">
                      <v-icon color="blue">mdi-thumb-down</v-icon>
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-content>
      </v-app>
    </div>
    <script src="/static/js/voting.js?v={{ .UnixTimestamp }}"></script>
  </body>
</html>
