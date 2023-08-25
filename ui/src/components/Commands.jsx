import {useState, React} from 'react'
import Table from 'react-bootstrap/Table';
import Badge from 'react-bootstrap/Badge';
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner'
import Form from 'react-bootstrap/Form';
import FloatingLabel from 'react-bootstrap/FloatingLabel';
import Button from 'react-bootstrap/Button';
import useSWR, { preload } from 'swr'


function Commands({setToastText}) {
  const [newCommand, setNewCommand] = useState("")
  const [newType, setNewType] = useState("Alias")
  const [newValue, setNewValue] = useState("")
  const fetcher = url => fetch(url).then(res => res.json())
  preload("/liveconfig", fetcher)
  const deleteCommand = (command) => {
    fetch(`/delete/${command}`, { "method": "DELETE" }).then(res => res.json())
      .then(data => {
        setToastText(`Deleted ${command}`)
        mutate(data)
      }
      )
  }
  const addCommand = () => {
    let url = encodeURIComponent(newValue)
    fetch(`/add/${newCommand}/${newType}?url=${url}`, { "method": "PUT" }).then(res => res.json())
      .then(data => {
        setNewCommand("")
        setNewValue("")
        setToastText(`Added ${newCommand}`)
        mutate(data)
      }
      )
  }

  const { data, mutate, error, isLoading } = useSWR("/liveconfig", fetcher)
  if (error) return (<div>Error</div>)
  if (isLoading) {
     return (<Spinner animation="border" variant='primary' />)
  }



  return (
    <Container>

      <h3>Commands <Badge bg="danger">{data.Entries?.length}</Badge></h3>
      <Table striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>Command</th>
            <th>Mode</th>
            <th>URL Template</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr key="add">
            <td>
              <FloatingLabel controlId="floatingCommand" label="Command">
                <Form.Control inputMode="text" placeholder="Command" value={newCommand} onChange={(e) => setNewCommand(e.currentTarget.value)} />
              </FloatingLabel>
            </td>
            <td>
              <FloatingLabel controlId="floatingSelect" label="Mode">
                <Form.Select aria-label="Mode" onChange={(e) => setNewType(e.currentTarget.value)}>
                  <option value="Alias">Alias</option>
                  <option value="Redirect">Redirect</option>
                  <option value="RedirectVarArgs">VarArgs</option>
                </Form.Select>
              </FloatingLabel>
            </td>
            <td>
              <FloatingLabel controlId="floatingValue" label="URL Template">
                <Form.Control type="text" inputMode="url" value={newValue} placeholder="URL" onChange={(e) => setNewValue(e.currentTarget.value)} />
              </FloatingLabel>
            </td>
            <td>
              <Button variant="primary" type="button" onClick={() => { addCommand() }}>Add</Button>
            </td>
          </tr>
          {data.Entries?.map((item, idx) => (
            <tr key={idx}>
              <td>{item.Command}</td>
              <td>{item.Type}</td>
              <td>{item.Value}</td>
              <td><Button variant="danger outline-warning" type="button" onClick={() => deleteCommand(item.Command)}>Delete</Button></td>
            </tr>
          ))}
        </tbody>
      </Table>
    </Container>
  )
}

export default Commands