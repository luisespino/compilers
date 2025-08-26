#include "symbol.h"
#include <stdio.h>
#include <search.h>

static int is_table_created = 0;

int symbol_table_create(size_t size) {
    if (is_table_created) {
        fprintf(stderr, "La tabla hash ya fue creada\n");
        return 0;
    }
    
    if (hcreate(size) == 0) {
        perror("Error al crear la tabla hash");
        return 0;
    }
    
    is_table_created = 1;
    return 1;
}

void symbol_table_destroy(void) {
    if (is_table_created) {
        hdestroy();
        is_table_created = 0;
    }
}

int symbol_table_insert(const char* name, int value) {
    if (!is_table_created) {
        fprintf(stderr, "La tabla hash no ha sido creada\n");
        return 0;
    }
    
    Symbol* new_entry = malloc(sizeof(Symbol));
    if (!new_entry) {
        perror("Error al asignar memoria");
        return 0;
    }
    
    new_entry->name = strdup(name);
    if (!new_entry->name) {
        perror("Error al copiar el nombre");
        free(new_entry);
        return 0;
    }
    
    new_entry->value = value;
    
    ENTRY e;
    e.key = new_entry->name;
    e.data = new_entry;
    /*
    ENTRY* result = hsearch(e, FIND);
    if (result != NULL) {
        fprintf(stderr, "La clave '%s' ya existe\n", name);
        free(new_entry->name);
        free(new_entry);
        return 0;
    }*/
    
    ENTRY* result = hsearch(e, ENTER);
    if (result == NULL) {
        fprintf(stderr, "Error al insertar en la tabla hash\n");
        free(new_entry->name);
        free(new_entry);
        return 0;
    }
    
    return 1;
}

int symbol_table_search(const char* name, int* value) {
    if (!is_table_created) {
        fprintf(stderr, "La tabla hash no ha sido creada\n");
        return 0;
    }
    
    ENTRY e;
    e.key = (char*)name;
    
    ENTRY* result = hsearch(e, FIND);
    if (result == NULL) {
        return 0;
    }
    
    Symbol* entry = (Symbol*)result->data;
    *value = entry->value;
    return 1;
}