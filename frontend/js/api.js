const api = {
    getHeaders: () => {
        const token = localStorage.getItem('token');
        return {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        };
    },

    get: async (url) => {
        const response = await fetch(url, {
            method: 'GET',
            headers: api.getHeaders()
        });
        return await response.json();
    },

    post: async (url, data) => {
        const response = await fetch(url, {
            method: 'POST',
            headers: api.getHeaders(),
            body: JSON.stringify(data)
        });
        return await response.json();
    },

    put: async (url, data) => {
        const response = await fetch(url, {
            method: 'PUT',
            headers: api.getHeaders(),
            body: JSON.stringify(data)
        });
        return await response.json();
    },

    delete: async (url) => {
        const response = await fetch(url, {
            method: 'DELETE',
            headers: api.getHeaders()
        });
        return await response.json();
    }
};