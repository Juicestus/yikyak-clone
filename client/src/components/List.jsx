import logo from "../logo.svg";
import Post from "./Post.jsx";

class List extends Component {
  update = (e) => {
    
  };

  constructor(props) {
    setInterval(update, 1000);
  }

  state = { seen: false };

  togglePop = () => {
    this.setState({
      seen: !this.state.seen,
    });
  };

  render() {
    return (
      <body>

      </body>
    );
  }
}

export default List;