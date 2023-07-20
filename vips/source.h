#include <vips/vips.h>

extern gint64 goSourceRead(char * owner_object_id, void *buffer, gint64 length);
extern gint64 goSourceSeek(char * owner_object_id, gint64 offset, int whence);

typedef struct _GoSourceArguments {
	char *owner_object_id;
} GoSourceArguments;

GoSourceArguments * create_go_source_arguments( char * owner_object_id );
VipsSourceCustom * create_go_custom_source( GoSourceArguments * source_args );
gulong connect_go_signal_read(VipsSourceCustom *source_custom, GoSourceArguments * source_args);
gulong connect_go_signal_seek(VipsSourceCustom *source_custom, GoSourceArguments * source_args);

void free_go_custom_source(VipsSourceCustom *source_custom, GoSourceArguments *source_args, gulong rsig_handler_id, gulong ssig_handler_id);


static gint64 go_read ( VipsSourceCustom *source_custom, void *buffer, gint64 length, GoSourceArguments * source_args );
static gint64 go_seek ( VipsSourceCustom *source_custom, gint64 offset, int whence, GoSourceArguments * source_args );
