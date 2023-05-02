import './App.css';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import {Table} from 'react-bootstrap';
import logoSvg from './img/boostchicken.svg';
import { useEffect, useState } from 'react';
 function App() {

  const [entries, setEntries] = useState({Entries:[]})
  useEffect(() => {
    const url = "/liveconfig";

    const fetchData = async () => {
      try {
        const response = await fetch(url);
        const json = await response.json();
        console.log(json);
        setEntries(json)
      } catch (error) {
        console.log("error", error);
      }
    };

    fetchData();
}, []);
  return (
    <div className="App">
      <header>
        <picture>
          <img className="logo" width={750} height={750} src={logoSvg} alt="What is a boostchicken?"/>
        </picture>
        </header>
      <main className="commands">
        <Table className="commands" size="sm" striped hover bordered variant="dark">
      <thead>
        <tr>
          <th>Command</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        {entries.Entries.map(data => (
              <tr>
                <td>{data.Command}</td>
                <td>{data.Value}</td>
              </tr>
            ))}

      </tbody>
    </Table>
      </main>
    </div>
  );
}

export default App;
