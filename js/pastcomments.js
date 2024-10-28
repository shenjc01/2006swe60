document.addEventListener('DOMContentLoaded', () => {
    const commentsContainer = document.getElementById('comments');

    fetch('/getComments')
        .then(res => res.json())
        .then(data => {
            data.comments.forEach(comment => {
                const commentDiv = document.createElement('div');
                commentDiv.className = 'comment';
                
                const username = document.createElement('div');
                username.className = 'username';
                username.textContent = comment.user.username;
                
                const text = document.createElement('div');
                text.className = 'text';
                text.textContent = comment.body;
                
                commentDiv.appendChild(username);
                commentDiv.appendChild(text);
                commentsContainer.appendChild(commentDiv);
            });
        })
        .catch(err => console.error('Error fetching comments:', err));
});
