import { Search } from "lucide-react";

const Navbar = () => {
  return (
    <div className="border-b border-gray-200 flex items-center justify-between max-lg:hidden px-6 mb-6">
      <div className="flex items-center justify-between">
        <div className="w-150 bg-white rounded-lg h-11 border border-gray-100 mx-8 my-4 shadow-sm flex items-center relative">
          <input
            type="text"
            className="text-gray-600 font-book-author mx-12 w-full outline-none"
            placeholder="Search books.."
          />
          <Search className="w-5 h-5 text-muted absolute left-4 text-gray-400" />
        </div>
      </div>
      <div className="w-10 h-10 rounded-full bg-bar text-white flex items-center justify-center font-bold font-book-author">
        <p>PJ</p>
      </div>
    </div>
  );
};

export default Navbar;
