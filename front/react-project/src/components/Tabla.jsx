import axios from "axios"
import { useEffect, useState } from "react"

const Tabla = () => {
    const [acciones, setAcciones] = useState([])
    const [itemsPerPage, setItemsPerPage] = useState(10)
    const [currentPage, setCurrentPage] = useState(1)
    const [filtro, setFiltro] = useState('')
    const [sortKey, setSortKey] = useState('')
    const [sortAsc, setSortAsc] = useState(true)

    useEffect(
        () => {
            axios.get('http://localhost:9000/api/datos')
            .then(response => setAcciones(response.data))
            .catch(error => console.error(error))
        }, []
    )
    const filteredItems = () => {
        let result = acciones

        if(filtro.length>0){
            result = result.filter((item) =>
            [item.ticker, item.company, item.brokerage].some(field =>
            field.toLowerCase().includes(filtro.toLowerCase())
            ))
        }

        if (sortKey.length>0) {
            if (["target_from", "target_to"].includes(sortKey)) {
                const key = sortKey;
                result = [...result].sort((a, b) => {
                const aVal = a[key];
                const bVal = b[key];
                return sortAsc ? aVal - bVal : bVal - aVal;
                });
            }
        }

        return result
    }

    const start = (currentPage - 1) * itemsPerPage
    const last = Math.ceil(acciones.length / itemsPerPage)
    const paginatedItems = filteredItems().slice(start, start + itemsPerPage)
    
    const handlefiltro = (event) => {
        setFiltro(event.target.value)
        
    }
    

    const handleItems = (event) => {
        setItemsPerPage(Number(event.target.value))
        setCurrentPage(1)
    }

    const restarPagina = () => {
        if(currentPage === 1){
            console.log('Primera página')
        } else {
            setCurrentPage(currentPage-1)
        }
    }

    const sumarPagina = () => {
        if(currentPage === last){
            console.log('Ultima página')
        } else {
            setCurrentPage(currentPage+1)
        }
    }
    
    const targetFrom = () => {
        setSortKey('target_from')
        if (sortAsc){
            setSortAsc(false)
        } else {
            setSortAsc(true)
        }
    }

    const targetTo = () => {
        setSortKey('target_to')
        if (sortAsc){
            setSortAsc(false)
        } else {
            setSortAsc(true)
        }
    }
    return(
        <>
        <div className="max-w-7xl mx-auto px-4 py-8">
            <h1 className="text-2xl font-bold mb-4 text-gray-800">Listado de Acciones</h1>

            <div className="mb-4 flex items-center justify-between">
                <input
                value={filtro}
                onChange={handlefiltro}
                type="text"
                placeholder="Buscar por ticker, empresa o brokerage..."
                className="w-full md:w-1/3 px-4 py-2 border rounded shadow-sm focus:outline-none focus:ring"
                />
                <div className="ml-4">
                <label className="text-sm font-medium mr-2">Por página:</label>
                <select className="border px-2 py-1 rounded" onChange={handleItems}>
                    <option value={10}>10</option>
                    <option value={25}>25</option>
                    <option value={50}>50</option>
                    <option value={100}>100</option>
                </select>
                </div>
            </div>

            {acciones.length > 0 ?
            <div>
                <table className="min-w-full divide-y divide-gray-200 rounded shadow text-sm">
                <thead className="bg-gray-100 sticky top-0 z-10">
                    <tr className="text-left text-gray-700">
                    <th className="px-3 py-2">Ticker</th>
                    <th className="px-3 py-2">Company</th>
                    <th className="px-3 py-2">Brokerage</th>
                    <th className="px-3 py-2">Action</th>
                    <th className="px-3 py-2">Rating from</th>
                    <th className="px-3 py-2">Rating to</th>
                    <th className="px-3 py-2 cursor-pointer" onClick={targetFrom}>
                        {<span>{sortKey === 'target_from' ? (sortAsc ? '↑' : '↓') : ''}</span>}
                        Target from
                    </th>
                    <th className="px-3 py-2 cursor-pointer" onClick={targetTo}>
                        {<span>{sortKey === 'target_to' ? (sortAsc ? '↑' : '↓') : ''}</span>}
                        Target to
                    </th>
                    </tr>
                </thead>
                <tbody className="divide-y divide-gray-100">
    
                    {paginatedItems.map(accion => (
                        <tr className="hover:bg-gray-50" key={accion.ID}>
                            <td className="px-3 py-2">{accion.ticker}</td>
                            <td className="px-3 py-2">{accion.company}</td>
                            <td className="px-3 py-2">{accion.brokerage}</td>
                            <td className="px-3 py-2">{accion.action}</td>
                            <td className="px-3 py-2">{accion.rating_from}</td>
                            <td className="px-3 py-2">{accion.rating_to}</td>
                            <td className="px-3 py-2">{accion.target_from}</td>
                            <td className="px-3 py-2">{accion.target_to}</td>
                        </tr> 
                    ))}
                    
                </tbody>
                </table>

                <div className="mt-4 flex items-center justify-between">
                    <button className="px-3 py-1 border rounded disabled:opacity-50"
                    onClick={restarPagina} disabled={currentPage === 1}>
                        Anterior
                    </button>
                    <span className="text-sm">Página {currentPage} de {last}</span>
                    <button className="px-3 py-1 border rounded disabled:opacity-50"
                    onClick={sumarPagina} disabled={currentPage === last}>
                        Siguiente
                    </button>
                </div>
            </div>
            :
            <div className="text-center text-gray-600">Cargando...</div>
            }
            

            
        </div>
        </>
    )
}

export default Tabla