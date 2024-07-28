const axios = require('axios');

const GRAFANA_URL = 'http://localhost:3000'; // replace with your Grafana URL if different
// const API_KEY = 'glsa_mBc9m8ZrwUk9MwAdWh2XusB6nO2jhMQr_94535c47'; 
const DASHBOARD_UID = 'PTSqcpJWk'; // dashboard UID from the URL

// Function to fetch dashboard JSON from Grafana
async function fetchDashboardJson() {
    try {
        const response = await axios.get(`${GRAFANA_URL}/api/dashboards/uid/${DASHBOARD_UID}`, {
            headers: {
                'Authorization': `Bearer ${API_KEY}`
            }
        });

        console.log('Dashboard JSON:', response.data);
        // Extract panel IDs from the dashboard JSON
        const panels = response.data.dashboard.panels;
        panels.forEach(panel => {
            console.log(`Panel Title: ${panel.title}, Panel ID: ${panel.id}`);
        });
    } catch (error) {
        console.error('Error fetching dashboard JSON from Grafana:', error);
    }
}

const PANEL_ID = 13; // replace with the actual panel ID

// Function to fetch panel data from Grafana
async function fetchPanelData() {
    try {
        // Example of fetching panel data
        const response = await axios.post(`${GRAFANA_URL}/api/ds/query`, {
            queries: [
                {
                    panelId: PANEL_ID,
                    dashboardId: DASHBOARD_UID,
                    range: {
                        from: 'now-1h',
                        to: 'now'
                    },
                    interval: '1m',
                    format: 'json'
                }
            ]
        }, {
            headers: {
                'Authorization': `Bearer ${API_KEY}`,
                'Content-Type': 'application/json'
            }
        });

        console.log('Panel Data:', response.data);
    } catch (error) {
        console.error('Error fetching panel data from Grafana:', error);
    }
}



// fetchDashboardJson();
fetchPanelData();

