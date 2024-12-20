new Vue({
    el: '#app',
    vuetify: new Vuetify(),
    data: {
      favorites: [], // Array to store favorite images
      view: 'grid',  // Current view ('grid' or 'list')
    },
    created() {
      this.fetchFavorites(); // Fetch favorites on page load
    },
    methods: {
      // Fetch favorite images from the API
      async fetchFavorites() {
        try {
            const response = await axios.get('/api/favorites', {
                headers: {
                    'x-api-key': 'live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg',
                },
            });// Replace with your actual API endpoint
          this.favorites = response.data.map((fav) => ({
            id: fav.id,
            url: fav.image.url,
          }));
        } catch (error) {
          console.error('Error fetching favorites:', error);
        }
      },
  
      // Toggle between grid and list views
      toggleView(viewType) {
        this.view = viewType;
      },
  
      // Navigate to other pages
      navigateTo(path) {
        window.location.href = path;
      },
    },
  });
  