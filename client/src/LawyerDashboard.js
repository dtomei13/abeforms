import React, { Component } from "react";
import axios from "axios";
import './App.css';
import { Card, Container, Image, Button} from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class LawyerDashboard extends Component{
    constructor(props){
        super(props);
        this.state = {
            case: "",
            items:[]

        }
    }
    componentDidMount() {
        this.getCase();
    }

    getCase = () => {
        axios.get(endpoint + "/lawyerdashboard/api/getcase").then(res => {
            console.log(res);
            if (res.data){
                console.log("setting info")
                this.setState({
                    items: res.data.map(item => {
                        
                        return(

                            <Card>
                            <Card.Content>
                                <Image
                                floated='right'
                                size='mini'
                                src='https://react.semantic-ui.com/images/avatar/large/molly.png'
                                />
                                <Card.Header>{item.firstname + " " + item.lastname}</Card.Header>
                                <Card.Meta>New Case</Card.Meta>
                                <Card.Description>
                               {item.description}
                                </Card.Description>
                            </Card.Content>
                            <Card.Content extra>
                                <div className='ui two buttons'>
                                <Button basic color='green' onClick = {() => this.caseComplete(item._id)}>
                                    Approve
                                </Button>
                                <Button basic color='red'>
                                    Decline
                                </Button>
                                </div>
                            </Card.Content>
                            </Card>
                                    );
                                })

                });
                console.log(this.setState)
            } else{
                this.setState({
                    items:[]
                });
            }
        })
    }

    caseComplete = id => {
        axios.put(endpoint + "/lawyerdashboard/api/takecase/" + id,{
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
              }
        })
    }
    render(){
        
        return (
            
        <Container>
        <div className="App">

        <h1>Open Cases</h1>
        <div className = "row">
            <Card.Group>{this.state.items}</Card.Group>
        </div>
        </div>
        </Container>
        )
    }

}

export default LawyerDashboard