import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import logoSvg from './img/boostchicken.svg';
import { useEffect, useState } from 'react';
import { Toast,Button, FloatingLabel, Form, Table } from 'react-bootstrap'
 function App() {

  const [showToast, setShowToast]= useState(false)
  const toggleToast = () => setShowToast(!showToast)
  const [toastText, setToastText] = useState("")
  const [history, setHistory] = useState([])
  const [entries, setEntries] = useState({Entries:[]})
  const [newCommand, setNewCommand] = useState("")
  const [newType, setNewType] = useState("")
  const [newValue, setNewValue] = useState("")

  const fetchData = async () => {
    try {
      const response = await fetch("/liveconfig");
      const json = await response.json();
      setEntries(json)

      const response2 = await fetch("/history");
      const json2 = await response2.json();
      setHistory(json2)
    } catch (error) {
      setToastText("Failed to query config")
    }
  };

  useEffect(() => {
    fetchData();
  }, []); 

  const deleteCommand = async (command) => {
    fetch("/delete/" + command, {method: 'DELETE'})
    .then(() => {
      setToastText("Deleted " + command)
      setShowToast(true)
    })
    .catch(setToastText("Failed to delete " + command))
    .finally(fetchData)
  }

  const addCommand = async () => {

    fetch("/add/"+newCommand+"/"+newType+"?url="+encodeURIComponent(newValue), {"method": "PUT"}).then(() => {
      fetchData()
      setToastText("Added " + newCommand)
    }).catch(() => {
      setToastText("Failed to add " + newCommand)
     
    }).finally(() => {
      setShowToast(true)
    })
  }
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
          <th>Type</th>
          <th>Value</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        {entries.Entries.map(data => (
              <tr>
                <td>{data.Command}</td>
                <td>{data.Type}</td>
                <td>{data.Value}</td>
                <td><Button as="button" type="button" onClick={() => deleteCommand(data.Command)}>Delete</Button></td>
              </tr>
            ))}
            <tr>

            <td>
              <FloatingLabel controlId="floatingCommand" label="Command">
                <Form.Control type="text" onChange={(e) => setNewCommand(e.currentTarget.value)} placeholder="Command" />
              </FloatingLabel>
          </td>
          <td>
              <Form.Select aria-label="Command Mode" onChange={(e) => setNewType(e.currentTarget.value)}>
              <option>Select one</option>
              <option value="Redirect">Redirect</option>
              <option value="RedirectVarArgs">RedirectVarArgs</option>
              <option value="Alias">Alias</option>
            </Form.Select>
          </td>
            <td>
            <FloatingLabel controlId="floatingValue" label="URL Template">
              <Form.Control type="text" onChange={(e) => setNewValue(e.currentTarget.value)} placeholder="https://www.google.com/?q=%s" />
            </FloatingLabel>
          </td>
          <td>
            <Button  as="button" type="button" onClick={() => addCommand()}>Add</Button>
          </td>
            
          </tr>
    
      </tbody>
    </Table>
    <Table className="commands" size="sm" striped hover bordered variant="dark">
      <thead>
        <tr>
          <th>Command</th>
          <th>Result</th>
          <th>IP</th>
        </tr>
      </thead>
      <tbody>
        {history.map(data => (
              <tr>
                <td>{data.Command}</td>
                <td>{data.Result}</td>
                <td>{data.IpAddress}</td>
              </tr>
            ))}
            </tbody></Table>
      </main>
      <Toast show={showToast}  onClose={toggleToast}>
      <Toast.Header>
        <strong className="me-auto">Boostchicken LOLww</strong>
      </Toast.Header>
      <Toast.Body>{toastText}</Toast.Body>
    </Toast>

    </div>

  );
}

export default App;
