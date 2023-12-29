"use client";
import Table from "react-bootstrap/Table";
import Badge from "react-bootstrap/Badge";
import Container from "react-bootstrap/Container";
import {useGetHistory} from "@boostchicken/lol-api";

function History() {
  const { data: history } = useGetHistory();
  return (
    <Container>
      <h3 className="input-group-text" style={{display: 'block'}}>
        History <Badge bg="primary">{history?.length}</Badge>
      </h3>
      <Table responsive striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>Command</th>
            <th>Result</th>
          </tr>
        </thead>
        <tbody>
          {history?.map((item, idx) => (
            <tr key={idx}>
              <td>{item.Command}</td>
              <td>
                <a href={item.Result}>{item.Result}</a>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </Container>
  );
}

export default History;
