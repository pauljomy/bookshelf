import logo from "@/assets/img/logo.svg";

const Header = () => {
  return (
    <div className="font-logo h-15  bg-sidebar  text-xl font-bold mb-4 shadow-sm">
      <div className="mx-5 flex items-center justify-left gap-3">
        <div>
          <img src={logo} className="w-8 h-14" />
        </div>
        <div>Bookshelf</div>
      </div>
    </div>
  );
};

export default Header;
