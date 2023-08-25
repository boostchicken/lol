import React from 'react'
import Table from 'react-bootstrap/Table';
import Badge from 'react-bootstrap/Badge';
import Container from 'react-bootstrap/Container';
import Spinner from 'react-bootstrap/Spinner'
import useSWR, {preload} from 'swr'


function History(props) {    
    const fetcher = url => fetch(url).then(res => res.json())
    preload("/history", fetcher)
    const { data, error, isLoading } = useSWR("/history", fetcher, { refreshInterval: 5000 })
    if(error) return(<div>Error</div>)
    if(isLoading) return(<Spinner animation="border"  variant='primary'/>)
    return(
        <Container>
        <h3>History <Badge bg="primary">{data.length}</Badge></h3>
        <Table striped bordered hover variant="dark">
        <thead>
            <tr>
            <th>Command</th>
            <th>Result</th>
            </tr>
        </thead>
        <tbody>
            {data.map((item, idx) => (
                <tr key={idx}>
                    <td>{item.Command}</td>
                    <td><a href={item.Result}>{item.Result}</a></td>
                </tr>
            ))}
        </tbody>
        </Table>
        </Container>
    )
}

export default History