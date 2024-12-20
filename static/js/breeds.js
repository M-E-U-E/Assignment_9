new Vue({
    el: '#app',
    vuetify: new Vuetify(),
    data: {
      breeds: [], // List of breeds
      selectedBreed: {}, // Currently selected breed
      images: [], // Images of the selected breed
      carouselIndex: 0, // Active carousel index
    },
    created() {
      this.fetchBreeds(); // Fetch the list of breeds on page load
    },
    watch: {
      images() {
        // Start carousel auto-play whenever images are updated
        this.startCarouselAutoPlay();
      },
    },
    methods: {
      // Fetch the list of breeds
      async fetchBreeds() {
        try {
          const response = await axios.get('https://api.thecatapi.com/v1/breeds', {
            headers: {
              'x-api-key': 'live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg',
            },
          });
  
          this.breeds = response.data;
          this.selectedBreed = this.breeds[0]; // Default to the first breed
          this.fetchBreedImages();
        } catch (error) {
          console.error('Error fetching breeds:', error);
        }
      },
  
      // Fetch images for the selected breed
      async fetchBreedImages() {
        try {
          const response = await axios.get('https://api.thecatapi.com/v1/images/search', {
            headers: {
              'x-api-key': 'live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg',
            },
            params: {
              breed_ids: this.selectedBreed.id,
              limit: 8, // Number of images to fetch
            },
          });
  
          this.images = response.data;
        } catch (error) {
          console.error('Error fetching breed images:', error);
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
  