
package vtgostack


type Entry interface{}



type Stack interface {
    Push(e Entry)
    Pop() Entry
    Top() Entry
    Depth() int
    RemCapacity() int
    Clear()
}

/* Adding flipping capability using interface embedding */
type FlippingStack interface {
    Stack
    Flip() FlippingStack
}



/* Implementation of stack using array */
type ArrayStack struct {
    data []Entry
    top int
}

// Creates new ArrayStack with specified maximum depth (capacity)
func NewArrayStack(maxDepth int) *ArrayStack {
    return &ArrayStack{ make([]Entry, maxDepth + 1), 0}
}


// Methods

func (s *ArrayStack) Push(e Entry) {
    if s.top < len(s.data) - 1 {
        s.top++
        s.data[s.top] = e
    }
}

func (s *ArrayStack) Top() Entry {
    if s.top > 0 {
        return s.data[s.top]
    }
    return nil
}


func (s *ArrayStack) Pop() Entry {
    var retVal Entry = nil
    if s.top > 0 {
        retVal = s.data[s.top]
        s.top--
    }
    return retVal
}

func (s *ArrayStack) Depth() int {
    return s.top
}

func (s *ArrayStack) RemCapacity() int {
    return len(s.data) - s.top - 1
}

func (s *ArrayStack) Clear() {
    s.top = 0
}

func (s *ArrayStack) MaxDepth() int {
    return len(s.data) - 1
}



/* Implementation of FlippingStack */
type FlippingArrayStack struct {
    *ArrayStack // Embedded
}

func (s *FlippingArrayStack) Flip() FlippingStack {
    flipped := NewArrayStack(s.MaxDepth())
    for s.Depth() > 0 {
        flipped.Push(s.Pop())
    }
    return &FlippingArrayStack{ flipped }
}



/* Stack implementation using linked list */

type element struct {
    next *element
    value Entry
}

type LinkedListStack struct {
    top *element
    maxDepth int
    depth int
}

// Create new LinkedListStack with specified maximum depth 
func NewLinkedListStack(maxDepth int) *LinkedListStack {
    return &LinkedListStack{ nil, maxDepth, 0 }
}

func (s *LinkedListStack) Push(e Entry) {
    if s.top == nil {
        s.top = &element{ nil, e}
        s.depth++
    } else if (s.depth < s.maxDepth) {
        newTop := &element{ s.top, e }
        s.top = newTop
        s.depth++
    } 
}


func (s *LinkedListStack) Top() Entry {
    if s.top == nil {
        return nil
    }
    return s.top.value
}

func (s *LinkedListStack) Pop() Entry {
    if s.top == nil {
        return nil
    }
    retVal := s.top.value
    s.top = s.top.next
    s.depth--
    return retVal
}

func (s *LinkedListStack) Depth() int {
    return s.depth
}

func (s *LinkedListStack) RemCapacity() int {
    return s.maxDepth - s.depth
} 

func (s *LinkedListStack) Clear() {
    s.top = nil
    s.depth = 0
}

func (s *LinkedListStack) MaxDepth() int {
    return s.maxDepth
}






type FlippingLinkedListStack struct {
    *LinkedListStack // Embedded
}


func (s *FlippingLinkedListStack) Flip() FlippingStack {
    flipped := NewLinkedListStack(s.MaxDepth())
    for s.Depth() > 0 {
        flipped.Push(s.Pop())
    }
    return &FlippingLinkedListStack{ flipped }
}


