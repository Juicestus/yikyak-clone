import React, { Component } from "react";

class PostPopup extends Component {
  state = { seen: false, message: "" };

  exit = () => this.props.toggle();
  
  createPost = (e) => {
    e.preventDefault();
    
    // Post the message!
    console.log(this.state.message);

    // Todo: charater count
    this.exit();
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
