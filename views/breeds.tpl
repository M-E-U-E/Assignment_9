<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Breeds</title>
    <link href="https://fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900|Material+Icons" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" rel="stylesheet">
    <link href="https://cdn.jsdelivr.net/npm/vuetify@2.x/dist/vuetify.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/breeds.css?v={{ .UnixTimestamp }}">
    <script src="https://cdn.jsdelivr.net/npm/vue@2.x/dist/vue.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/vuetify@2.x/dist/vuetify.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <script src="/static/js/breeds.js?v=1.0"></script>
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
                <v-btn @click="navigateTo('/breeds')" class="nav-btn active">
                  <v-icon left>mdi-magnify</v-icon> Breeds
                </v-btn>
                <v-btn @click="navigateTo('/favs')" class="nav-btn">
                  <v-icon left>mdi-heart</v-icon> Favs
                </v-btn>
              </v-col>
            </v-row>

            <!-- Breeds Dropdown -->
            <v-row justify="center" class="mt-2">
              <v-col cols="12" md="6">
                <v-select
                  :items="breeds"
                  item-text="name"
                  v-model="selectedBreed"
                  label="Breeds"
                  return-object
                  attach
                  @change="onBreedChange"
                ></v-select>
              </v-col>
            </v-row>

            <!-- Breed Card -->
            <v-row justify="center" class="mt-2">
              <v-col cols="12" md="6">
                <v-card>
                  <!-- Breed Carousel -->
                  <v-carousel v-model="carouselIndex">
                    <v-carousel-item
                      v-for="(image, i) in images"
                      :key="i"
                      :src="image.url"
                    ></v-carousel-item>
                  </v-carousel>

                  <!-- Breed Information -->
                  <v-card-title>
                    <h3>{{ "{{ selectedBreed.name }}" }}</h3>
                  </v-card-title>
                  <v-card-text>
                    <p>{{ "{{ selectedBreed.description }}" }}</p>
                    <p><strong>Temperament:</strong> {{ "{{ selectedBreed.temperament }}" }}</p>
                    <p><strong>Origin:</strong> {{ "{{ selectedBreed.origin }}" }}</p>
                    <v-btn :href="selectedBreed.wikipedia_url" target="_blank" color="orange" flat>Wikipedia</v-btn>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-content>
      </v-app>
    </div>
    <script src="/static/js/breeds.js?v={{ .UnixTimestamp }}"></script>
  </body>
</html>
