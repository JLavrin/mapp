import {useCallback, useEffect, useState} from 'react'
import { GoogleMap, Marker, useJsApiLoader } from '@react-google-maps/api'
import Circle from './icons/Circle';
import { faCircle } from '@fortawesome/free-solid-svg-icons';

const containerStyle = {
  width: '100vw',
  height: '100vh'
};

const center = {
  lat: 54.373459,
  lng: 18.620453
};

type Vehicle = {
  routeShortName: string
  vehicleId: number
  lat: number
  lon: number
}

function App() {
  const { isLoaded } = useJsApiLoader({
    id: 'google-map-script',
    googleMapsApiKey: import.meta.env.GOOGLE_API_KEY
  })

  const [vehicles, setVehicles] = useState<Vehicle[]>([])
  const [map, setMap] = useState(null)

  useEffect(() => {
    const ws = new WebSocket('/ws')
    console.log(ws)

    ws.onmessage = (event) => {
      const parsed = JSON.parse(event.data)
      switch (parsed.event) {
        case 'vehicles_update':
          setVehicles(parsed.data.vehicles)
          console.log(parsed.data.vehicles)
          break
        case 'connected':
          break
        default:
          break
      }

    }

    return () => {
      ws.close()
    }
  }, [])

  const onLoad = useCallback((map) => {
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
      zoom={11}
      onLoad={onLoad}
      onUnmount={onUnmount}
      options={{
        disableDoubleClickZoom: true,
        minZoom: 10,
        maxZoom: 14,
      }}
    >
      {vehicles.map((vehicle) => (
        <Marker icon={{
          // @ts-ignore
          path: faCircle.icon[4],
          strokeWeight: 1,
          strokeColor: 'black',
          scale: 0.055,
        }} key={vehicle.vehicleId} position={{ lat: vehicle.lat, lng: vehicle.lon }} title={vehicle.routeShortName}/>
      ))}
    </GoogleMap>
  ) : <></>
}

export default App
