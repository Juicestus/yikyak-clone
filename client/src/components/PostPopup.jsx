import React, { Component } from "react";
import axios from "axios";

const endpoint = () => window.location.hostname;

// this will be stored in a cookie given durring auth
const DEMO_USER_CODE = "b14ec87f9690454035f81493666d57fd"; // randomly generated hex

class PostPopup extends Component {
  state = { seen: false, message: "" };

  exit = () => this.props.toggle();
  
  createPost = (e) => {
    e.preventDefault();
    
    if (!this.state.message.replace(/\s/g, '').length) {
      alert("Please enter a valid message.");
      return;
    }

    axios.post('/api/post_message', {
      message: this.state.message,
      usercode: DEMO_USER_CODE 
    }, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      }
    }).then(res => {
      this.exit();
      console.log(res);
    }).catch(err => {
      console.log(err);
    });

    // Todo: charater count
  };

  render() {
    return (
      <div className="popup-struct">
        <div className="popup-content">
          <span className="popup-close" onClick={this.exit}>&times;{" "}</span>
          <h2>Create a Post</h2>
          <form id="post-msg-form" onSubmit={this.createPost}>
            <textarea onChange={e => this.setState({message: e.target.value}) }
              className="post-text"
              placeholder="What's on your mind?"
              value={this.state.message}
            ></textarea>
              <a className="transparent-button" onClick={this.createPost}>Create Post</a>
          </form>
        </div>
      </div>
    );
  }
}

export default PostPopup;
