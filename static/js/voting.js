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
                const response = await axios.get('https://api.thecatapi.com/v1/images/search', {
                    headers: {
                        'x-api-key': 'live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg',
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
            const response = await axios.post('/api/favorites', payload);
            console.log('Added to favorites:', response.data);
            alert('Image added to favorites!');
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
  