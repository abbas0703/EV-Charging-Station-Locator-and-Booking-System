let map;
let selectedStation;

function initMap() {
    const mapCenter = { lat: 37.7749, lng: -122.4194 }; // Default to San Francisco

    map = new google.maps.Map(document.getElementById("map"), {
        center: mapCenter,
        zoom: 12,
    });

    fetchChargingStations(mapCenter);
}

function fetchChargingStations(location) {
    fetch(`/charging-stations?location=${location.lat},${location.lng}&radius=5000`)
        .then(response => response.json())
        .then(data => {
            const stations = data.stations;
            stations.forEach(station => {
                const marker = new google.maps.Marker({
                    position: {
                        lat: station.geometry.location.lat,
                        lng: station.geometry.location.lng,
                    },
                    map: map,
                    title: station.name,
                });

                marker.addListener("click", () => {
                    selectedStation = station;
                    displayStationInfo(station);
                });
            });
        });
}

function displayStationInfo(station) {
    document.getElementById("station-name").innerText = station.name;
}

document.getElementById("book-slot").addEventListener("click", () => {
    if (!selectedStation) {
        alert("Please select a station first.");
        return;
    }

    const bookingData = {
        user_id: "user123", // Replace with actual user ID
        charging_station_id: selectedStation.place_id,
        start_time: "2024-08-22T10:00:00Z",
        end_time: "2024-08-22T11:00:00Z",
    };

    fetch("/book-slot", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(bookingData),
    })
        .then(response => response.json())
        .then(data => {
            alert(data.status);
        });
});
