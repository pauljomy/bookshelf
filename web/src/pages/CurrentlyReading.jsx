import SingleCard from "../components/singleCard";

const CurrentlyReading = ({ bookList }) => {
  return (
    <div className=" ">
      <div className="mx-5 font-logo flex justify-between items-baseline mb-5">
        <h1 className="text-2xl font-bold">Currently Reading</h1>
        <p className="text-[0.8 5em]">{bookList.length} books</p>
      </div>
      <div>
        {bookList.map((book) => (
          <SingleCard key={book.id} book={book} />
        ))}
      </div>
    </div>
  );
};

export default CurrentlyReading;
