import { useState, useEffect } from "react";
import Sidebar from "./components/sidebar";
import Navbar from "./components/navbar";
import AllBooks from "./pages/AllBooks";

const App = () => {
    const [bookList, setBookList] = useState([]);

    useEffect(() => {
        const fetchBookList = async () => {
            const response = await fetch("http://localhost:4040/v1/books");
            const data = await response.json();
            setBookList(data.books);
        };

        fetchBookList();
    }, []);

    console.log(bookList);

    return (
        <main className="min-h-screen">
            <div className="lg:flex min-h-screen">
                <Sidebar bookList={bookList} />
                <div className="flex flex-col flex-1">
                    <Navbar />
                    <AllBooks bookList={bookList} />
                </div>
            </div>
        </main>
    );
};

export default App;
