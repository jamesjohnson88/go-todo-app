function Header() {
    return (
        <header className="bg-gray-200 shadow-sm">
            <div className="container mx-auto px-4">
                <div className="flex items-center justify-center h-16">
                    <h1 className="text-3xl sm:text-4xl font-bold text-gray-800 text-center" style={{
                        fontFamily: '"Coming Soon", cursive',
                        textDecoration: 'underline',
                        textDecorationStyle: 'double'
                    }}>
                        my never-ending to-do list
                    </h1>
                </div>
            </div>
            <style>{`
                @import url('https://fonts.googleapis.com/css2?family=Coming+Soon&display=swap');
            `}</style>
        </header>
    )
}

export default Header