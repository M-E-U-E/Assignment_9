new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data: {
    breeds: [],             // List of breeds
    selectedBreed: {},      // Currently selected breed
    images: [],             // Images of the selected breed
    carouselIndex: 0,       // Active carousel index
    apiKey: '',             // API key fetched dynamically
  },
  async created() {
    await this.fetchApiKey(); // Fetch the API key on page load
    await this.fetchBreeds(); // Fetch the list of breeds after retrieving the API key
  },
  watch: {
    images() {
      // Start carousel auto-play whenever images are updated
      this.startCarouselAutoPlay();
    },
  },
  methods: {
    // Fetch the API key dynamically from the server
    async fetchApiKey() {
      try {
        const response = await axios.get('/api/config/api_key'); // Endpoint to get the API key
        this.apiKey = response.data.api_key;
        console.log('API Key Fetched:', this.apiKey);
      } catch (error) {
        console.error('Error fetching API key:', error.message);
        alert('Failed to load API key. Some features may not work.');
      }
    },

    // Fetch the list of breeds
    async fetchBreeds() {
      if (!this.apiKey) {
        console.warn('API key is missing. Cannot fetch breeds.');
        return;
      }

      try {
        const response = await axios.get('https://api.thecatapi.com/v1/breeds', {
          headers: {
            'x-api-key': this.apiKey, // Use dynamically fetched API key
          },
        });

        this.breeds = response.data;
        if (this.breeds.length > 0) {
          this.selectedBreed = this.breeds[0]; // Default to the first breed
          this.fetchBreedImages();
        }
      } catch (error) {
        console.error('Error fetching breeds:', error.message);
        alert('Failed to load breeds.');
      }
    },

    // Fetch images for the selected breed
    async fetchBreedImages() {
      if (!this.apiKey) {
        console.warn('API key is missing. Cannot fetch breed images.');
        return;
      }

      try {
        const response = await axios.get('https://api.thecatapi.com/v1/images/search', {
          headers: {
            'x-api-key': this.apiKey, // Use dynamically fetched API key
          },
          params: {
            breed_ids: this.selectedBreed.id,
            limit: 8, // Number of images to fetch
          },
        });

        this.images = response.data;
      } catch (error) {
        console.error('Error fetching breed images:', error.message);
        alert('Failed to load breed images.');
      }
    },

    // Handle breed selection change
    onBreedChange() {
      this.fetchBreedImages();
    },

    // Navigate to different pages
    navigateTo(path) {
      window.location.href = path;
    },

    // Start automatic carousel changes
    startCarouselAutoPlay() {
      setInterval(() => {
        if (this.images.length > 0) {
          this.carouselIndex = (this.carouselIndex + 1) % this.images.length;
        }
      }, 3000); // Change every 3 seconds
    },
  },
});
