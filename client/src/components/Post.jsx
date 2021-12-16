import logo from "../logo.svg";

function Post(props) {
  return (
    <article>
        <h3>{props.author}</h3>
        <p>{props.body}</p>
    </article>
  );
}

export default Post;