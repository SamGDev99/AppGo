const Titulo = () => {
    return(
        <>
        <header className="flex items-center justify-center gap-3 px-6 py-4 mb-6 rounded-2xl bg-white shadow animate-fade-in-up">
            <svg
            xmlns="http://www.w3.org/2000/svg"
            className="w-10 h-10 text-blue-600"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            strokeWidth="2"
            >
            <path strokeLinecap="round" strokeLinejoin="round" d="M3 3v18h18" />
            <path strokeLinecap="round" strokeLinejoin="round" d="M9 9l3 3-3 3" />
            </svg>
            <h1 className="text-4xl font-bold text-gray-800">Informe de Acciones</h1>
        </header>
        </>
    )
}

export default Titulo