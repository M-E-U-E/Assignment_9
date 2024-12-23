new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data: {
    activeTab: 'favs',         // Indicates the current active tab
    favourites: [],            // Array to store favorite images
    limit: 12,                 // Number of images per page
    page: 1,                   // Current page index (starts at 1)
    pagination_count: 0,       // Total count of favorites for pagination
    viewMode: 'gallery grid',  // Current view mode (grid or list)
  },
  created() {
    this.getFavourites(); // Fetch favorite images when the component is created
  },
  methods: {
    // Fetch favorite images dynamically
    async getFavourites() {
      try {
        let query_params = {
          limit: this.limit,   // Number of items per page
          order: 'DESC',       // Order of results (newest first)
          page: this.page - 1, // API expects zero-based page index
        };

        // Make GET request to The Cat API
        let response = await axios.get('https://api.thecatapi.com/v1/favourites', {
          params: query_params,
          headers: {
            'x-api-key': 'live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg', // API key
          },
        });

        // Update favourites and pagination data
        console.log('Fetched Favorites:', response.data);
        this.favourites = response.data;
        this.pagination_count = response.headers['pagination-count']; // Total count of favorites

      } catch (err) {
        console.error('Error fetching favorites:', err.message);
        alert('Failed to load favorites.');
      }
    },

    // Change the view mode (grid or list)
    setView(view) {
      this.viewMode = view === 'grid' ? 'gallery grid' : 'gallery list';
    },

    // Navigate to different pages
    navigateTo(path) {
      window.location.href = path;
    },

    // Change the page for pagination
    async changePage(newPage) {
      this.page = newPage; // Update the current page
      await this.getFavourites(); // Fetch new data for the updated page
    },
  },
});
