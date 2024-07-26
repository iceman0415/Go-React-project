import React from "react";
import "./Person.css";

class Person extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
          checked: localStorage.getItem(props.email)
        };
        this.handleChecked = this.handleChecked.bind(this);
    };

    handleChecked() {
        if (this.getState() === "") {
            this.setState({checked: "checked"});
        } else {
            this.setState({checked: ""});
        }
    };

    getState() {
        return this.state.checked;
    };

    componentDidUpdate(prevProps, prevState) {
        localStorage.setItem(this.props.email, this.state.checked);
    };

    render() { 
        return (
            <div className="person">
                <span>{this.props.name}, {this.props.age}, {this.props.balance}, {this.props.email}, {this.props.address}</span><input type="checkbox" checked={this.state.checked} onChange={this.handleChecked}></input>
            </div>
        )   
    };
}

export default Person;