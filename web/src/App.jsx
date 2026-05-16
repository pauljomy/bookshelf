import { useState, useEffect } from "react";
import Header from "./components/header";
import CurrentlyReading from "@/pages/CurrentlyReading";

const App = () => {
  const [bookList, setBookList] = useState([]);

  useEffect(() => {
    const fetchBookList = async () => {
      const response = await fetch("/data.json");
      const data = await response.json();
      setBookList(data);
    };

    fetchBookList();
  }, []);

  console.log(bookList);

  return (
    <main>
      <Header />
      <CurrentlyReading bookList={bookList} />
    </main>
  );
};

export default App;
