import { useEffect, useState } from 'react'
import axios from 'axios'
import './App.css'
import Titulo from './components/Titulo'
import Tarjeta from './components/Tarjeta'
import Tabla from './components/Tabla'

const App = () => {
  const [accion, setAccion] = useState([])

  useEffect(
    () => {
      axios.get('http://localhost:9000/calcularDatos')
      .then(response => setAccion(response.data))
      .catch(error => console.error(error))
    }, [])
  
  return(
    <>
      <div className="min-h-screen bg-gray-50">
        <div className="max-w-6xl mx-auto px-4 py-8 space-y-8">
          <Titulo />
          {accion.length > 0 ?
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Tarjeta accion={accion[0]} tipo={'mejor'}/>
              <Tarjeta accion={accion[1]} tipo={'peor'}/>
            </div>
            :
            <p>Cargando datos...</p>
          }
          <Tabla />
        </div>
      </div>
    </>
  ) 

  
}

export default App
