const SingleCard = ({ book }) => {
    return (
        <div className=" bg-white px-4 py-4 flex gap-3 rounded-md border shadow-sm border-gray-100 cursor-pointer">
            <div>
                <img
                    className="rounded-md overflow-hidden w-30 h-45"
                    src={book.cover}
                />
            </div>
            <div className="flex flex-col flex-1">
                <h3 className="text-[1.1em] font-medium font-logo text-book-tile leading-6 mb-1  ">
                    {book.title}
                </h3>
                <p className="font-book-author text-sm text-book-author mb-2 ">
                    {book.author}
                </p>
                <div className="bg-genre  text-[0.75em] text-genre-text font-bold rounded-full w-fit py-1 px-3 text-center">
                    <p className=" ">{book.genre}</p>
                </div>
                <div className="mt-auto ">
                    <div className="flex justify-between items-center w-full mb-1">
                        <p className="text-[0.75em] text-genre-text">
                            Progress
                        </p>
                        <p className="text-[0.75em] text-bar font-bold">
                            {book.progress}%
                        </p>
                    </div>
                    <div className="h-1.5 bg-gray-100 w-full rounded-full">
                        <div
                            className="h-full bg-bar rounded-full"
                            style={{ width: `${book.progress}%` }}
                        ></div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default SingleCard;
