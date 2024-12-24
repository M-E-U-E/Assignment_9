new Vue({
    el: '#app',
    vuetify: new Vuetify(),
    data: {
      activeTab: 'voting',
      catImage: '',
      catImageId: '', // Store the current image ID
      validExtensions: ['jpg', 'png', 'jpeg', 'gif'], // Valid image extensions
    },
    created() {
      this.fetchCatImage(); // Fetch the initial random image
    },
    methods: {
      // Fetch a random image with valid extensions
      async fetchCatImage() {
        try {
          // Fetch the API key from a server-side endpoint
          const apiKeyResponse = await axios.get('/api/config/api_key');
          const apiKey = apiKeyResponse.data.api_key;
  
          // Fetch the cat image using the API key
          const response = await axios.get('https://api.thecatapi.com/v1/images/search', {
            headers: {
              'x-api-key': apiKey,
            },
          });
  
          if (response.data.length > 0) {
            const validImages = response.data.filter((item) => {
              const extension = item.url.split('.').pop().toLowerCase();
              return this.validExtensions.includes(extension);
            });
  
            if (validImages.length > 0) {
              this.catImage = validImages[0].url;
              this.catImageId = validImages[0].id;
            } else {
              this.catImage = '/static/images/fallback.jpg';
            }
          } else {
            console.warn('No images returned by API.');
            this.catImage = '/static/images/fallback.jpg';
          }
        } catch (error) {
          console.error('Error fetching cat image:', error.message);
          this.catImage = '/static/images/fallback.jpg';
        }
      },
  
      // Add the current image to favorites
      async addToFavorites() {
        try {
          const payload = { image_id: this.catImageId };
  
          // Send the favorite image data to the server
          const response = await axios.post('/api/favorites', payload);
          console.log('Added to favorites:', response.data);
        //   alert('Image added to favorites!');
  
          // Fetch a new random image after successfully adding to favorites
          await this.fetchCatImage();
        } catch (error) {
          console.error('Error adding to favorites:', error.message);
          alert('Failed to add to favorites.');
        }
      },
  
      // Navigate to different pages
      navigateTo(path) {
        window.location.href = path;
      },
    },
  });
  