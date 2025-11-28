import os
import json
import uuid
from datetime import datetime
from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS
from pathlib import Path
from dotenv import load_dotenv

# 加载环境变量
load_dotenv()

app = Flask(__name__, static_folder='dist')
CORS(app)

# 数据文件路径
DATA_DIR = Path(__file__).parent / 'data'
DATA_FILE = DATA_DIR / 'todos.json'

def ensure_data_file():
    """确保数据目录和文件存在"""
    DATA_DIR.mkdir(exist_ok=True)
    if not DATA_FILE.exists():
        DATA_FILE.write_text('[]', encoding='utf-8')

def read_todos():
    """读取所有 Todo"""
    try:
        with open(DATA_FILE, 'r', encoding='utf-8') as f:
            data = json.load(f)
            return data if isinstance(data, list) else []
    except Exception:
        return []

def write_todos(todos):
    """写入 Todo 数据"""
    tmp_file = DATA_FILE.with_suffix('.tmp')
    with open(tmp_file, 'w', encoding='utf-8') as f:
        json.dump(todos, f, ensure_ascii=False, indent=2)
    tmp_file.replace(DATA_FILE)

def is_valid_iso_date(date_str):
    """验证 ISO 日期格式"""
    if not isinstance(date_str, str):
        return False
    try:
        datetime.fromisoformat(date_str.replace('Z', '+00:00'))
        return True
    except Exception:
        return False

# 初始化数据文件
ensure_data_file()

@app.route('/api/todos', methods=['GET'])
def get_todos():
    """获取所有 Todo"""
    todos = read_todos()
    return jsonify({'data': todos})

@app.route('/api/todos', methods=['POST'])
def create_todo():
    """创建新 Todo"""
    body = request.get_json() or {}
    title = (body.get('title') or '').strip()
    due_at = body.get('dueAt')
    
    if not title:
        return jsonify({'error': '标题不能为空'}), 400
    
    if due_at and not is_valid_iso_date(due_at):
        return jsonify({'error': '时间格式不正确'}), 400
    
    todos = read_todos()
    now = datetime.now().isoformat()
    
    new_todo = {
        'id': str(uuid.uuid4()),
        'title': title,
        'dueAt': due_at,
        'completed': False,
        'createdAt': now,
        'updatedAt': now
    }
    
    todos.append(new_todo)
    write_todos(todos)
    
    return jsonify({'data': new_todo}), 201

@app.route('/api/todos/<todo_id>', methods=['PUT'])
def update_todo(todo_id):
    """更新 Todo"""
    body = request.get_json() or {}
    todos = read_todos()
    
    todo_index = next((i for i, t in enumerate(todos) if t['id'] == todo_id), -1)
    if todo_index == -1:
        return jsonify({'error': '未找到Todo'}), 404
    
    todo = todos[todo_index].copy()
    
    if 'title' in body:
        title = (body['title'] or '').strip()
        if not title:
            return jsonify({'error': '标题不能为空'}), 400
        todo['title'] = title
    
    if 'dueAt' in body:
        due_at = body['dueAt']
        if due_at is not None and not is_valid_iso_date(due_at):
            return jsonify({'error': '时间格式不正确'}), 400
        todo['dueAt'] = due_at
    
    if 'completed' in body and isinstance(body['completed'], bool):
        todo['completed'] = body['completed']
    
    todo['updatedAt'] = datetime.now().isoformat()
    todos[todo_index] = todo
    write_todos(todos)
    
    return jsonify({'data': todo})

@app.route('/api/todos/<todo_id>/toggle', methods=['POST'])
def toggle_todo(todo_id):
    """切换 Todo 完成状态"""
    todos = read_todos()
    
    todo_index = next((i for i, t in enumerate(todos) if t['id'] == todo_id), -1)
    if todo_index == -1:
        return jsonify({'error': '未找到Todo'}), 404
    
    todo = todos[todo_index].copy()
    todo['completed'] = not todo['completed']
    todo['updatedAt'] = datetime.now().isoformat()
    
    todos[todo_index] = todo
    write_todos(todos)
    
    return jsonify({'data': todo})

@app.route('/api/todos/<todo_id>', methods=['DELETE'])
def delete_todo(todo_id):
    """删除 Todo"""
    todos = read_todos()
    
    todo_index = next((i for i, t in enumerate(todos) if t['id'] == todo_id), -1)
    if todo_index == -1:
        return jsonify({'error': '未找到Todo'}), 404
    
    removed = todos.pop(todo_index)
    write_todos(todos)
    
    return jsonify({'data': removed})

# 前端路由支持
@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def serve(path):
    """提供前端静态文件"""
    if path and (Path(app.static_folder) / path).exists():
        return send_from_directory(app.static_folder, path)
    return send_from_directory(app.static_folder, 'index.html')

if __name__ == '__main__':
    port = int(os.getenv('PORT', 5000))
    app.run(host='0.0.0.0', port=port, debug=True)
