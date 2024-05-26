import { useCallback, useState } from 'react'
import { GoogleMap, Marker, useJsApiLoader } from '@react-google-maps/api'

const containerStyle = {
  width: '400px',
  height: '400px'
};

const center = {
  lat: 54.373459,
  lng: 18.620453
};

function App() {
  const { isLoaded } = useJsApiLoader({
    id: 'google-map-script',
    googleMapsApiKey: import.meta.env.GOOGLE_API_KEY
  })

  const [map, setMap] = useState(null)

  const onLoad = useCallback((map) => {
    // This is just an example of getting and using the map instance!!! don't just blindly copy!
    // @ts-ignore
    const bounds = new window.google.maps.LatLngBounds(center);
    map.fitBounds(bounds);

    setMap(map)
  }, [])

  const onUnmount = useCallback((map) => {
    setMap(null)
  }, [])

  return isLoaded ? (
    <GoogleMap
      mapContainerStyle={containerStyle}
      center={center}
      zoom={11.6}
      onLoad={onLoad}
      onUnmount={onUnmount}
      options={{
        disableDefaultUI: true,
        disableDoubleClickZoom: true,
        minZoom: 11.6,
        maxZoom: 11.6,
      }}
    >
      <Marker position={center} />
    </GoogleMap>
  ) : <></>
}

export default App
