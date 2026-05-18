import SingleCard from "../components/singleCard";

const AllBooks = ({ bookList }) => {
    return (
        <div className="lg:px-6 ">
            <div className="mx-5 font-logo flex max-md:justify-between items-baseline mb-5 md:gap-5">
                <h1 className="text-2xl font-bold lg:text-3xl">All Books</h1>
                <p className="text-[0.85em] font-book-author text-book-author">
                    {bookList.length} books
                </p>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-5 px-5  xl:grid-cols-3">
                {bookList.map((book) => (
                    <SingleCard key={book.id} book={book} />
                ))}
            </div>
        </div>
    );
};

export default AllBooks;
