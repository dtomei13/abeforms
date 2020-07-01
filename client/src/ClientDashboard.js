import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import {
    Card,
    Container,
    Image,
    Button,
    Modal,
    Header,
    ModalActions,
} from "semantic-ui-react";
import { Redirect } from "react-router-dom";
let endpoint = "http://localhost:8080";
var userEmail;
var clientEmail;

class ClientDashboard extends Component {
    constructor(props) {
        super(props);
        this.state = {
            case: "",
            assigned_items: [],
            unassigned_items: [],
            redirect: "",
            lawyerEmail: "",
            selectedTime: "2020-06-24T22:00:00Z",
            clientEmail: "",
        };
    }
    componentDidMount() {
        this.getAssignedCase();
        this.getUnassignedCase();
    }

    getAssignedCase = () => {
        axios.get(endpoint + "/clientdashboard/api/assignedcases",{headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },}).then((res) => {
            //console.log(res);
            if (res.data) {
                //console.log("setting open cases info");
                this.setState({
                    assigned_items: res.data.map((item) => {
                        return (
                            <Card>
                                <Card.Content>
                                    <Image
                                        floated="right"
                                        size="mini"
                                        src="https://react.semantic-ui.com/images/avatar/large/molly.png"
                                    />
                                    <Card.Header>
                                        {item.firstname + " " + item.lastname}
                                    </Card.Header>
                                    <Card.Meta>{item.stateofissue}</Card.Meta>
                                    <Card.Description>{item.description}</Card.Description>
                                </Card.Content>
                                <Card.Content extra>
                                    <Modal
                                        trigger={
                                            <Button basic color="green" content="Green">
                                                Details
                                            </Button>
                                        }
                                    >
                                        <Modal.Header>Some Details Go here</Modal.Header>
                                        <Modal.Content image>
                                            <Image
                                                wrapped
                                                size="medium"
                                                src="https://react.semantic-ui.com/images/avatar/large/rachel.png"
                                            />
                                            <Modal.Description>
                                                <Header>Client Details</Header>
                                                <p>
                                                    Name: {item.firstname} {item.lastname}
                                                    <br></br>
                                                    Description: {item.description} <br></br>
                                                    Location: {item.stateofissue} <br></br>
                                                    Available Times: Times go here...
                                                </p>
                                            </Modal.Description>
                                        </Modal.Content>
                                        <Modal.Actions>
                                            <Button basic color="red">
                                                Close
                                            </Button>
                                        </Modal.Actions>
                                    </Modal>
                                </Card.Content>
                            </Card>
                        );
                    }),
                });
                console.log(this.setState);
            } else {
                this.setState({
                    assigned_items: [],
                });
            }
        });
    };

    getUnassignedCase = () => {
        axios.get(endpoint + "/clientdashboard/api/unassignedcases").then((res) => {
            if (res.data) {
                //console.log("setting my cases info");
                this.setState({
                    unassigned_items: res.data.map((item) => {
                        //console.log("hell111o");
                        //console.log(item);
                        return (
                            <Card>
                                <Card.Content>
                                    <Image
                                        floated="right"
                                        size="mini"
                                        src="https://react.semantic-ui.com/images/avatar/large/molly.png"
                                    />
                                    <Card.Header>
                                        {item.firstname + " " + item.lastname}
                                    </Card.Header>
                                    <Card.Meta>{item.stateofissue}</Card.Meta>
                                    <Card.Description>{item.description}</Card.Description>
                                </Card.Content>
                                <Card.Content extra>
                                    <div className="ui buttons">
                                        <Modal
                                            trigger={
                                                <Button basic color="violet" content="Violet">
                                                    Details
                                                </Button>
                                            }
                                        >
                                            <Modal.Header>Some Details Go here</Modal.Header>
                                            <Modal.Content image>
                                                <Image
                                                    wrapped
                                                    size="medium"
                                                    src="https://react.semantic-ui.com/images/avatar/large/rachel.png"
                                                />
                                                <Modal.Description>
                                                    <Header>Client Details</Header>
                                                    Name: {item.firstname} {item.lastname}
                                                    <br></br>
                                                    Description: {item.description} <br></br>
                                                    Location: {item.stateofissue} <br></br>
                                                    Available Times: Times go here...
                                                </Modal.Description>
                                            </Modal.Content>
                                            <ModalActions>
                                            </ModalActions>
                                        </Modal>
                                    </div>
                                </Card.Content>
                            </Card>
                        );
                    }),
                });
                console.log(this.setState);
                //console.log(this.setState)
            } else {
                this.setState({
                    unassigned_items: [],
                });
            }
        });
    };

    // this.setState({userEmail: this.props.location.state.email})
    render() {
        try {
            const { lawyerEmail, selectedTime, clientEmail } = this.state;
            console.log(this.state);
            console.log("HERE");
        } catch (e) {
            return <Redirect to={"/clientdashboard/sign_in"} />; //Check if user is authenticated
        }

        return (
            <Container>
                <div className="App">
                    <h1>Unassigned Cases</h1>
                    <div className="row">
                        <Card.Group>{this.state.unassigned_items}</Card.Group>
                    </div>

                    <h1>Assigned Cases</h1>
                    <div className="row">
                        <Card.Group>{this.state.assigned_items}</Card.Group>
                    </div>
                </div>
            </Container>
        );
    }
}
export default ClientDashboard;
