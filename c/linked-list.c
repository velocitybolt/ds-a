#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

struct node {
  int data;
  struct node* next;
};

// Iterate thorough list until null and increment counter
int get_length(struct node* head) {
  int length = 0;
  while (head != NULL) {
    length++;
    head = head->next;
  }
  return length;
}

void printList(struct node* head) {
  while (head != NULL) {
    printf("This is the current value : %d", head->data);
    head = head->next;
  }
}

struct node* get_nth(struct node* head, int n) {
  if (n < 0) {
    return NULL;
  }
  int curr = 0;
  while (head != NULL) {
    // handles zero condition
    if (curr == n) {
      return head;
    }
    curr++;
    head = head->next;
  }
  return NULL;
}

void append_data(struct node** headRef, int data) {
  assert(headRef != NULL);
  if (*headRef == NULL) {
    *headRef = (struct node*)malloc(sizeof(struct node));
    (*headRef)->data = data;
    (*headRef)->next = NULL;
    return;
  }
  struct node* head = *headRef;
  while (head->next != NULL) {
    head = head->next;
  }
  head->next = (struct node*)malloc(sizeof(struct node));
  head->next->data = data;
  head->next->next = NULL;
}

int pop(struct node** headRef) {
  // make sure the pointer pointing to head isnt null, bc that is basically the
  // head ptr..
  assert(headRef != NULL);
  // if the actual head is not null ...
  if (*headRef != NULL) {
    // dereference the headRef to get actual pointer to head
    struct node* head = *headRef;
    // save data inside head
    int data = head->next->data;
    // free memory assoc. w head
    free(head);
    // have head point to head.next which breaks link to current head..
    head = head->next;
    // return value inside of hed
    return data;
  }
  assert(0);
}

int main() {
  struct node* head = (struct node*)malloc(sizeof(struct node));
  head->data = 77;
}
