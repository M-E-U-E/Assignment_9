new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data: {
    favourites: [],
    apiKey: '',
    viewMode: 'grid', // Default view mode
  },
  async created() {
    await this.getApiKey();
    await this.getFavourites();
  },
  methods: {
    async getApiKey() {
      try {
        const response = await axios.get('/api/config/api_key');
        this.apiKey = response.data.api_key;
      } catch (err) {
        console.error('Error fetching API key:', err.message);
        alert('Failed to load API key. Some features may not work.');
      }
    },

    async getFavourites() {
      if (!this.apiKey) {
        console.warn('API key is missing. Cannot fetch favorites.');
        return;
      }

      try {
        const response = await axios.get('https://api.thecatapi.com/v1/favourites', {
          headers: {
            'x-api-key': this.apiKey,
          },
        });

        this.favourites = response.data;
        this.favourites.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
      } catch (err) {
        console.error('Error fetching favorites:', err.message);
        alert('Failed to load favorites.');
      }
    },

    navigateTo(path) {
      window.location.href = path;
    },
  },
});
