import Header from "./header";
import { BookMarked } from "lucide-react";

const Sidebar = ({ bookList }) => {
    return (
        <div className="lg:w-60 bg-sidebar">
            <Header />

            <div className="flex justify-between items-center cursor-pointer  hover:bg-sidebar-hover  hover:text-sidebar-text-hover h-10 py-3 px-5 text-book-author font-medium">
                <div className="flex gap-2">
                    <BookMarked className="w-6 h-6" />
                    <p className="">All books</p>
                </div>
                <p className="">{bookList.length}</p>
            </div>
        </div>
    );
};

export default Sidebar;
