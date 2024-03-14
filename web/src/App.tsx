import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import NavBar from "./components/navbar";
import TypographyDemo from "./components/title";

function App() {
  return (
    <>
      <div className="topnav font-imb-plex-mono container">
        <NavBar />
      </div>
      <div>
        <TypographyDemo />
      </div>
    </>
  );
}

export default App;
